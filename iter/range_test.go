package iter

import (
	"fmt"
	"github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleRange() {
	positive := func(i int) bool { return i > 0 }
	positives := Filter[int](Range(-10, 10, 2), positive)
	fmt.Println(Collect[int](positives))
	// Output: [2 4 6 8]
}

func TestRange(t *testing.T) {
	it := Range(0, 2, 1)
	it.Next()
	assert.Equal(t, it.Value(), 0)
	it.Next()
	assert.Equal(t, it.Value(), 1)
	assert.Assert(t, !it.Next())
}

func TestRangeEmpty(t *testing.T) {
	assert.Assert(t, !Range(0, 0, 1).Next())
}
