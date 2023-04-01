package iter

import (
	"fmt"
	"testing"

	"github.com/nnnewb/battery/internal/assert"
)

func ExampleChain() {
	fmt.Println(
		Collect(
			Chain(Lift([]int{1, 2}),
				Lift([]int{3, 4}),
				Lift([]int{0, 9}),
			)))
	// Output: [1 2 3 4 0 9]
}

func TestChainMultiple(t *testing.T) {
	it := Chain(Lift([]int{1, 2}), Lift([]int{3, 4}))

	it.Next()
	assert.Equal(t, it.Value(), 1)
	it.Next()
	assert.Equal(t, it.Value(), 2)
	it.Next()
	assert.Equal(t, it.Value(), 3)
	it.Next()
	assert.Equal(t, it.Value(), 4)
	assert.Assert(t, !it.Next())
}

func TestChainSingle(t *testing.T) {
	it := Chain(Lift([]int{1, 2}))

	it.Next()
	assert.Equal(t, it.Value(), 1)
	it.Next()
	assert.Equal(t, it.Value(), 2)
	assert.Assert(t, !it.Next())
}

func TestChainEmpty(t *testing.T) {
	assert.Assert(t, !Chain[int]().Next())
}

func TestChainExhausted(t *testing.T) {
	delegate1 := Exhausted[int]()
	delegate2 := Exhausted[int]()
	it := Chain(delegate1, delegate2)

	it.Next()
	assert.Assert(t, !it.Next())
	it.Next()
	assert.Assert(t, !it.Next())
}
