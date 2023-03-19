package iter_test

import (
	"fmt"
	"github.com/nnnewb/battery/assert"
	"testing"

	"github.com/nnnewb/battery/iter"
)

func ExampleChain() {
	fmt.Println(
		iter.Collect[int](
			iter.Chain[int](iter.Lift([]int{1, 2}),
				iter.Lift([]int{3, 4}),
				iter.Lift([]int{0, 9}),
			)))
	// Output: [1 2 3 4 0 9]
}

func TestChainMultiple(t *testing.T) {
	it := iter.Chain[int](iter.Lift([]int{1, 2}), iter.Lift([]int{3, 4}))

	it = it.Next()
	assert.Equal(t, it.Value(), 1)
	it = it.Next()
	assert.Equal(t, it.Value(), 2)
	it = it.Next()
	assert.Equal(t, it.Value(), 3)
	it = it.Next()
	assert.Equal(t, it.Value(), 4)
	it = it.Next()
	assert.Assert(t, it.Exhausted())
}

func TestChainSingle(t *testing.T) {
	it := iter.Chain[int](iter.Lift([]int{1, 2}))

	it = it.Next()
	assert.Equal(t, it.Value(), 1)
	it = it.Next()
	assert.Equal(t, it.Value(), 2)
	it = it.Next()
	assert.Assert(t, it.Exhausted())
}

func TestChainEmpty(t *testing.T) {
	assert.Assert(t, iter.Chain[int]().Next().Exhausted())
}

func TestChainExhausted(t *testing.T) {
	delegate1 := iter.Mock[int]()
	delegate2 := iter.Mock[int]()
	it := iter.Chain[int](delegate1, delegate2)

	it = it.Next()
	assert.Assert(t, it.Exhausted())
	it = it.Next()
	assert.Assert(t, it.Exhausted())

	assert.Equal(t, delegate1.NextCallCount(), 0)
	assert.Equal(t, delegate2.NextCallCount(), 0)
}
