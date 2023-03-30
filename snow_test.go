package snow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	snow := NewSnow("", 1)
	var map2 = map[uint64]bool{}
	for i := 0; i < 10000; i++ {
		map2[snow.Gen()] = true
	}
	assert.Equal(t, len(map2), 10000, "they should be equal")
}
