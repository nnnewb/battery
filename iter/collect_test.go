package iter

import (
	"testing"

	"github.com/nnnewb/battery/internal/assert"
)

func TestCollect(t *testing.T) {
	items := Collect(Range(0, 5, 1))
	assert.Equal(t, items, []int{0, 1, 2, 3, 4})
}

func TestCollectEmpty(t *testing.T) {
	items := Collect(Exhausted[int]())
	assert.Empty(t, items)
}
