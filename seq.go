package snow

import (
	"os"
	"sync"
	"syscall"
)

type Seq struct {
	i         uint64
	logCnt    uint64
	diskValue uint64
	name      string
	fp        *os.File
	mutex     sync.Mutex
}

func (seq *Seq) Init() uint64 {
	file := seq.name + ".txt"
	b := exists(file)
	if seq.logCnt == 0 {
		seq.logCnt = 256
	}
	if !b {
		fp, _ := os.Create(file)
		defer fp.Close()
		fp.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0})
	}
	seq.diskValue = seq.DiskGet()
	if seq.diskValue > seq.i {
		seq.i = seq.diskValue
	}
	return seq.i
}

var mSeq = map[string]*Seq{}
var lockSeq sync.Mutex

func NewSeq(key string) *Seq {
	seq, has := mSeq[key]
	if !has {
		lockSeq.Lock()
		defer lockSeq.Unlock()

		seq, has := mSeq[key]
		if has {
			return seq
		}
		seq2 := Seq{name: key}
		seq2.Init()
		mSeq[key] = &seq2
		return &seq2
	} else {
		return seq
	}
}

func (seq *Seq) Incr() uint64 {
	return seq.IncrBy(1)
}

func (seq *Seq) IncrBy(value uint64) uint64 {
	seq.mutex.Lock()
	seq.i += value
	if seq.i >= seq.diskValue {
		seq.diskValue = seq.i + seq.logCnt
		seq.DiskSet(seq.diskValue)
	}
	localI := seq.i
	seq.mutex.Unlock()
	return localI
}

func (seq *Seq) DiskIncr() uint64 {
	return seq.DiskIncrBy(1)
}

func (seq *Seq) DiskIncrBy(value uint64) uint64 {
	seq.mutex.Lock()
	seq.i += value
	seq.diskValue = seq.i
	seq.DiskSet(seq.diskValue)
	localI := seq.i
	seq.mutex.Unlock()
	return localI
}

func (seq *Seq) DiskSet(value uint64) uint64 {
	if seq.fp == nil {
		seq.fp, _ = os.OpenFile(seq.name+".txt", syscall.O_RDWR, 0666)
	}
	var off int64 = int64((value >> 16) << 3)
	seq.fp.WriteAt(uint64byte(value), off)
	return value
}

func (seq *Seq) DiskGet() uint64 {
	if seq.fp == nil {
		seq.fp, _ = os.OpenFile(seq.name+".txt", syscall.O_RDWR, 0666)
	}
	fi, _ := os.Stat(seq.name + ".txt")

	b := make([]byte, 8)

	if fi.Size() == 0 {
		seq.fp.ReadAt(b, 0)
	} else {
		seq.fp.ReadAt(b, fi.Size()-8)
	}
	j := byte2uint64(b)
	return j
}

func exists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func byte2uint64(b []byte) uint64 {
	len := uint64(len(b))
	var x uint64 = 0
	j := len - 1
	var i uint64 = 0
	for ; i < len; i++ {
		x |= uint64(b[j]) << (i * 8)
		j--
	}
	return x
}

func uint64byte(u uint64) []byte {
	var i uint64 = 7
	b := make([]byte, 8)
	var j uint64 = 0
	var k int = 0
	for {
		j = u / (1 << (i * 8))
		if j != 0 {
			u -= (j << (i * 8))
			b[k] = byte(j)
		}
		if i == 0 {
			break
		}
		k++
		i--
	}
	return b
}
