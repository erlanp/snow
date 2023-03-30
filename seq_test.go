package snow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeq(t *testing.T) {
	seq := NewSeq("seq")
	var map2 = map[uint64]bool{}
	for i := 0; i < 10000; i++ {
		map2[seq.Incr()] = true
	}
	assert.Equal(t, len(map2), 10000, "they should be equal")
}
