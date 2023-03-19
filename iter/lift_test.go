package iter

import (
	"fmt"
	"github.com/nnnewb/battery/assert"
	"testing"
)

func ExampleLift() {
	positives := Filter[int](Lift([]int{-1, 4, 6, 4, -5}), func(i int) bool { return i > -1 })
	fmt.Println(Collect[int](positives))
	// Output: [4 6 4]
}

func TestLift(t *testing.T) {
	it := Lift([]int{1, 2}).Next()
	assert.Equal(t, it.Value(), 1)
	it = it.Next()
	assert.Equal(t, it.Value(), 2)
	it = it.Next()
	assert.Assert(t, it.Exhausted())
}

func TestLiftEmpty(t *testing.T) {
	assert.Assert(t, Lift([]int{}).Next().Exhausted())
}
