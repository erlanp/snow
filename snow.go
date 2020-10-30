package snow

import (
    "time"
    crypto "crypto/rand"
    "sync"
)

type Snow struct {
    seq Seq
    i uint64
    i1 uint64
    i2 uint64
    shard_id uint64
    rand_i1 uint64
}

func (this *Snow) Init() uint64 {
    this.i = this.seq.Init();
    return this.i;
}

func (this *Snow) incr() uint64 {
    this.i = this.seq.Incr();
    return this.i;
}

var m_snow = map[string]*Snow{};
func NewSnow(key string) *Snow {
    snow2, has := m_snow[key];
    if (has != true) {
        seq := Seq{name:key+"snow.txt", log_cnt: 8};
        snow := Snow{seq:seq, shard_id:5}
        m_snow[key] = &snow;
        return &snow;
    } else {
        return snow2;
    }
}

func (this *Snow) Gen() uint64 {
    // t := uint64(time.Now().Unix() - 1500000000) >> 8;
    t := uint64(time.Now().Unix() - 1500000000);
    var mutex sync.Mutex
    mutex.Lock()
    if ((t % 2) == 1) {
        if (this.i1 == 0) {
            this.i2 = 0;
            this.rand_i1 = uint64(rand_byte());
        }
        this.i1++;
        if (this.i1 >= 256) {
            this.i1 = 0;
            this.i2 = 0;
            this.incr();
        }
        mutex.Unlock()
        return (t + this.i * 2) << 16 | (this.shard_id << 8) | this.i1 
    } else {
        if (this.i2 == 0) {
            this.i1 = 0;
        }
        this.i2++;
        if (this.i2 >= 256) {
            this.i2 = 0;
            this.i1 = 0;
            this.incr();
        }
        mutex.Unlock()
        return (t + this.i * 2) << 16 | (this.shard_id << 8) | this.i2
    }
}

func rand_byte() byte {
    b := []byte{0,};
    crypto.Read(b);
    return b[0];
}