package iter

import (
	"fmt"
	assert2 "github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleChain() {
	fmt.Println(
		Collect[int](
			Chain[int](Lift([]int{1, 2}),
				Lift([]int{3, 4}),
				Lift([]int{0, 9}),
			)))
	// Output: [1 2 3 4 0 9]
}

func TestChainMultiple(t *testing.T) {
	it := Chain[int](Lift([]int{1, 2}), Lift([]int{3, 4}))

	it = it.Next()
	assert2.Equal(t, it.Value(), 1)
	it = it.Next()
	assert2.Equal(t, it.Value(), 2)
	it = it.Next()
	assert2.Equal(t, it.Value(), 3)
	it = it.Next()
	assert2.Equal(t, it.Value(), 4)
	it = it.Next()
	assert2.Assert(t, it.Exhausted())
}

func TestChainSingle(t *testing.T) {
	it := Chain[int](Lift([]int{1, 2}))

	it = it.Next()
	assert2.Equal(t, it.Value(), 1)
	it = it.Next()
	assert2.Equal(t, it.Value(), 2)
	it = it.Next()
	assert2.Assert(t, it.Exhausted())
}

func TestChainEmpty(t *testing.T) {
	assert2.Assert(t, Chain[int]().Next().Exhausted())
}

func TestChainExhausted(t *testing.T) {
	delegate1 := Exhausted[int]()
	delegate2 := Exhausted[int]()
	it := Chain[int](delegate1, delegate2)

	it = it.Next()
	assert2.Assert(t, it.Exhausted())
	it = it.Next()
	assert2.Assert(t, it.Exhausted())
}
