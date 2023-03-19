package iter

import (
	"fmt"
	"github.com/nnnewb/battery/assert"
	"testing"
)

func ExampleTake() {
	it := Take[int](Count(), 2)

	it = it.Next()
	fmt.Println(it.Value())
	it = it.Next()
	fmt.Println(it.Value())
	it = it.Next()
	fmt.Println(it.Exhausted())
	// Output:
	// 0
	// 1
	// true
}

func TestTakeIter(t *testing.T) {
	it := Take[int](Count(), 2)

	it = it.Next()
	assert.Equal(t, it.Value(), 0)
	it = it.Next()
	assert.Equal(t, it.Value(), 1)
	it = it.Next()
	assert.Assert(t, it.Exhausted())
}

func TestTakeIterEmpty(t *testing.T) {
	it := Take[int](Count(), 0)
	it = it.Next()
	assert.Assert(t, it.Exhausted())
}

func TestTakeExhausted(t *testing.T) {
	delegate := Exhausted[int]()
	it := Take[int](delegate, 10)

	it = it.Next()
	assert.Assert(t, it.Exhausted())
	it = it.Next()
	assert.Assert(t, it.Exhausted())
}
