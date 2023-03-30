package snow

import (
	"sync"
	"time"
)

type Snow struct {
	seq     *Seq
	i       uint64
	i1      uint64
	i2      uint64
	shardId uint64
	mutex   sync.Mutex
}

func (snow *Snow) incr() uint64 {
	snow.i = snow.seq.Incr()
	return snow.i
}

var mSnow = map[string]*Snow{}
var lockSnow sync.Mutex

func NewSnow(key string, shardId uint64) *Snow {
	snow2, has := mSnow[key]
	if !has {
		lockSnow.Lock()
		defer lockSnow.Unlock()

		snow2, has := mSnow[key]
		if has {
			return snow2
		}
		seq := NewSeq(key + "_snow")
		snow := Snow{seq: seq, shardId: shardId}
		mSnow[key] = &snow
		return &snow
	} else {
		return snow2
	}
}

func (snow *Snow) Gen() uint64 {
	t := uint64(time.Now().Unix() - 1500000000)
	snow.mutex.Lock()
	defer snow.mutex.Unlock()
	if (t % 2) == 1 {
		if snow.i1 == 0 {
			snow.i2 = 0
		}
		snow.i1++
		if snow.i1 >= 256 {
			snow.i1 = 0
			snow.i2 = 0
			snow.incr()
		}
		return (t+snow.i*2)<<16 | (snow.shardId << 8) | snow.i1
	} else {
		if snow.i2 == 0 {
			snow.i1 = 0
		}
		snow.i2++
		if snow.i2 >= 256 {
			snow.i2 = 0
			snow.i1 = 0
			snow.incr()
		}
		return (t+snow.i*2)<<16 | (snow.shardId << 8) | snow.i2
	}
}
