package iter

import (
	"fmt"
	assert2 "github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleLift() {
	positives := Filter[int](Lift([]int{-1, 4, 6, 4, -5}), func(i int) bool { return i > -1 })
	fmt.Println(Collect[int](positives))
	// Output: [4 6 4]
}

func TestLift(t *testing.T) {
	it := Lift([]int{1, 2})
	it.Next()
	assert2.Equal(t, it.Value(), 1)
	it.Next()
	assert2.Equal(t, it.Value(), 2)
	assert2.Assert(t, !it.Next())
}

func TestLiftEmpty(t *testing.T) {
	assert2.Assert(t, !Lift([]int{}).Next())
}
