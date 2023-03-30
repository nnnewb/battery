package iter

import (
	"fmt"
	"github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleFilter() {
	it := Filter[int](Lift([]int{0, 1, 0, 2}), func(i int) bool { return i == 0 })
	for it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// 0
	// 0
}

func TestFilter(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	it := Filter[int](Range[int](0, 10, 1), isEven)
	it.Next()
	assert.Equal(t, it.Value(), 0)
	it.Next()
	assert.Equal(t, it.Value(), 2)
}

func TestFilterEmpty(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	it := Filter[int](Exhausted[int](), isEven)
	assert.Assert(t, !it.Next())
}

func TestFilterExhausted(t *testing.T) {
	delegate := Exhausted[int]()
	it := Filter[int](delegate, func(_ int) bool { return true })

	assert.Assert(t, !it.Next())
	assert.Assert(t, !it.Next())
}

func TestFilterExhaustedLater(t *testing.T) {
	delegate := Lift([]int{1})
	it := Filter[int](delegate, func(_ int) bool { return true })

	it.Next()
	assert.Equal(t, it.Value(), 1)
	assert.Assert(t, !it.Next())
}
