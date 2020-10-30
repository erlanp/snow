package snow

import (
    "os"
    "syscall"
    "time"
    "sync"
    // "fmt"
)

type Seq struct {
    i uint64
    log_cnt uint64
    disk_value uint64
    name string
    fp *os.File
    fp2 *os.File
    mutex *sync.Mutex
}

func (this *Seq) SetName(name string) {
    this.name = name;
}

func (this *Seq) Init() uint64 {
    file := this.name;
    b := exists(file);
    if (this.log_cnt == 0) {
        this.log_cnt = 32;
    }
    if (b == false) {
        fp, _ := os.Create(file);
        fp.Write([]byte{0,0,0,0,0,0,0,1});
        fp.Close();
    }
    this.disk_value = this.DiskGet();
    if (this.disk_value > this.i) {
        this.i = this.disk_value;
    }
    return this.i;
}

var m_seq = map[string]*Seq{};
func NewSeq(key string) *Seq {
    seq, has := m_seq[key];
    if (has != true) {
        var seq2 Seq;
        seq2.SetName(key+"seq.txt");
        seq2.Init();
        m_seq[key] = &seq2;
        return &seq2;
    } else {
        return seq;
    }
}

func unix_time() int64 {
    return time.Now().Unix();
}

func unix_time_nano() int64 {
    return time.Now().UnixNano();
}

func (this *Seq) Incr() uint64 {
    var mutex sync.Mutex
    mutex.Lock()
    this.i++;
    if (this.i >= this.disk_value) {
        this.disk_value = this.i + this.log_cnt
        this.DiskSet(this.disk_value);
    }
    mutex.Unlock()
    return this.i;
}

func (this *Seq) IncrBy(value uint64) uint64 {
    this.i += value;
    if (this.i >= this.disk_value) {
        this.disk_value = this.i + this.log_cnt
        this.DiskSet(this.disk_value);
    }
    
    return this.i;
}

func (this *Seq) DiskSet(value uint64) uint64 {
    if (this.fp == nil) {
        this.fp, _ = os.OpenFile(this.name, syscall.O_RDWR, 0666);
    }
    this.fp.WriteAt(uint64byte(value), 0);
    return value;
}

func (this *Seq) DiskGet() uint64 {
    if (this.fp == nil) {
        this.fp, _ = os.OpenFile(this.name, syscall.O_RDWR, 0666);
    }
    b := make([]byte, 8);
    this.fp.ReadAt(b, 0);
    j := byte2uint64(b);
    return j;
}

func (this *Seq) DiskIncr() uint64 {
    if (this.fp == nil) {
        this.fp, _ = os.OpenFile(this.name, syscall.O_RDWR, 0666);
    }
    b := make([]byte, 8);
    this.fp.ReadAt(b, 0);
    j := byte2uint64(b);
    j ++;
    this.fp.WriteAt(uint64byte(j), 0);
    return j;
}

func (this *Seq) DiskIncrBy(value uint64) uint64 {
    if (this.fp == nil) {
        this.fp, _ = os.OpenFile(this.name, syscall.O_RDWR, 0666);
    }
    b := make([]byte, 8);
    this.fp.ReadAt(b, 0);
    j := byte2uint64(b);
    j += value;
    this.fp.WriteAt(uint64byte(j), 0);
    return j;
}

func incr(file string) uint64 {
    var fp *os.File;
    if (exists(file)) {
        fp, _ = os.OpenFile(file, syscall.O_RDWR, 0666);
        defer fp.Close();
        b := make([]byte, 8);
        fp.Read(b);
        j := byte2uint64(b);
        j++;
        fp2, _ := os.Create(file);
        defer fp2.Close();
        fp2.Write(uint64byte(j));
        return j;
    } else {
        fp, _ = os.Create(file);
        defer fp.Close();
        fp.Write([]byte{0,0,0,0,0,0,0,1});
        return uint64(1);
    }
}

func exists(file string) bool {
    _, err := os.Stat(file); 
    return !os.IsNotExist(err);
}

func byte2uint64(b []byte) uint64 {

    len := uint64(len(b));
    var x uint64 = 0;
    j := len - 1;
    var i uint64 = 0
    for ; i < len; i++ {
        x |= uint64(b[j]) << (i*8);
        j--;
    }
    return x;
}

func uint64byte(u uint64) []byte {
    var i uint64 = 7;
    b := make([]byte, 8);
    var j uint64 = 0;
    var k int = 0;
    for ;; {
        j = u/(1 << (i*8));
        if (j != 0) {
            u -= (j << (i*8))
            b[k] = byte(j);
        }
        if (i == 0) {
            break;
        }
        k++;
        i--;
    }
    return b;
}