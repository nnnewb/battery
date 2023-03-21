package iter

import (
	"fmt"
	"github.com/nnnewb/battery/assert"
	"testing"
)

func ExampleRange() {
	positive := func(i int) bool { return i > 0 }
	positives := Filter[int](Range(-10, 10, 2), positive)
	fmt.Println(Collect[int](positives))
	// Output: [2 4 6 8 10]
}

func TestRange(t *testing.T) {
	it := Range(0, 2, 1).Next()
	assert.Equal(t, it.Value(), 1)
	it = it.Next()
	assert.Equal(t, it.Value(), 2)
	it = it.Next()
	assert.Assert(t, it.Exhausted())
}

func TestRangeEmpty(t *testing.T) {
	assert.Assert(t, Range(0, 0, 1).Next().Exhausted())
}
