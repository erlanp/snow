package snow

import (
    "testing"
    "math/rand"
)



var g uint64;
var g2 int64;
var bb []byte = make([]byte, 8);
var bb2 []byte;

func BenchmarkUnixTimeNano(b *testing.B) {
    // ns := []byte{1, 2, 3, 4, 5, 6, 7, 8}
    for i := 0; i < b.N; i++ {
        g2 = unix_time_nano()
    }
}

func BenchmarkSnow(b *testing.B) {
    snow := []Snow {
        {seq:Seq{name:"snow1.txt", log_cnt: 8}, shard_id:1,},
        {seq:Seq{name:"snow2.txt", log_cnt: 8}, shard_id:2,},
        {seq:Seq{name:"snow3.txt", log_cnt: 8}, shard_id:3,},
        {seq:Seq{name:"snow4.txt", log_cnt: 8}, shard_id:4,},
        {seq:Seq{name:"snow5.txt", log_cnt: 8}, shard_id:5,},
    };
    for _, element := range snow {
        element.Init();
    }
    for i := 0; i < b.N; i++ {
        snow[0].Gen();
        snow[1].Gen();
        snow[2].Gen();
        snow[3].Gen();
        snow[4].Gen();
    }
}

func BenchmarkRand(b *testing.B) {
    // ns := []byte{1, 2, 3, 4, 5, 6, 7, 8}
    for i := 0; i < b.N; i++ {
        rand.Intn(255)
    }
}

// func BenchmarkIncr(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         g = incr("test.txt")
//     }
// }

func BenchmarkIncr(b *testing.B) {
    co := Seq{i:1, name:"test2.txt",};
    co.Init();
    for i := 0; i < b.N; i++ {
        g = co.Incr()
    }
}

func BenchmarkIncrBy(b *testing.B) {
    co := Seq{i:1, name:"test3.txt",};
    co.Init();
    for i := 0; i < b.N; i++ {
        g = co.IncrBy(2)
    }
}

func BenchmarkDiskIncr(b *testing.B) {
    co := Seq{i:1, name:"test4.txt",};
    co.Init();
    for i := 0; i < b.N; i++ {
        g = co.DiskIncr()
    }
}

func BenchmarkDiskIncrBy(b *testing.B) {
    co := Seq{i:1, name:"test5.txt",};
    co.Init();
    for i := 0; i < b.N; i++ {
        g = co.DiskIncrBy(2)
    }
}





