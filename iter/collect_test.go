package iter

import (
	"github.com/nnnewb/battery/internal/assert"
	"testing"
)

func TestCollect(t *testing.T) {
	items := Collect[int](Range[int](0, 5, 1))
	assert.Equal(t, items, []int{0, 1, 2, 3, 4})
}

func TestCollectEmpty(t *testing.T) {
	items := Collect[int](Exhausted[int]())
	assert.Empty(t, items)
}
