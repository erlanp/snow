package snow

import (
	"testing"
)

var g uint64

func BenchmarkSnow(b *testing.B) {
	sequence := []*Seq{
		&Seq{name: "snow1"},
		&Seq{name: "snow2"},
		&Seq{name: "snow3"},
		&Seq{name: "snow4"},
		&Seq{name: "snow5"},
	}
	snow := []Snow{
		Snow{seq: sequence[0], shardId: 0},
		Snow{seq: sequence[1], shardId: 1},
		Snow{seq: sequence[2], shardId: 2},
		Snow{seq: sequence[3], shardId: 3},
		Snow{seq: sequence[4], shardId: 4},
	}
	for i := 0; i < b.N; i++ {
		snow[0].Gen()
		snow[1].Gen()
		snow[2].Gen()
		snow[3].Gen()
		snow[4].Gen()
	}
}
func BenchmarkSnowTwo(b *testing.B) {
	snow := NewSnow("snow", 5)
	for i := 0; i < b.N; i++ {
		g = snow.Gen()
	}
}

func BenchmarkIncr(b *testing.B) {
	seq := Seq{i: 1, name: "test2.txt"}
	seq.Init()
	for i := 0; i < b.N; i++ {
		g = seq.Incr()
	}
}

func BenchmarkIncrBy(b *testing.B) {
	seq := Seq{i: 1, name: "test3.txt"}
	seq.Init()
	for i := 0; i < b.N; i++ {
		g = seq.IncrBy(2)
	}
}

func BenchmarkDiskIncr(b *testing.B) {
	seq := Seq{i: 1, name: "test4.txt"}
	seq.Init()
	for i := 0; i < b.N; i++ {
		g = seq.DiskIncr()
	}
}

func BenchmarkDiskIncrBy(b *testing.B) {
	seq := Seq{i: 1, name: "test5.txt"}
	seq.Init()
	for i := 0; i < b.N; i++ {
		g = seq.DiskIncrBy(2)
	}
}
