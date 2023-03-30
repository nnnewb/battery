package iter

import (
	"fmt"
	"github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleSkip() {
	it := Skip[int](Range[int](0, 10, 1), 2)
	it.Next()
	fmt.Println(it.Value())
	// Output: 2
}

func TestSkip(t *testing.T) {
	counter := Skip[int](Range[int](0, 10, 1), 2)
	counter.Next()
	assert.Equal(t, counter.Value(), 2)
}

func TestSkipExhausted(t *testing.T) {
	it := Skip[int](Exhausted[int](), 5)
	assert.Assert(t, !it.Next())
	assert.Assert(t, !it.Next())
}

func TestSkipExhaustedLater(t *testing.T) {
	delegate := Lift([]int{42, 43})
	it := Skip[int](delegate, 1)
	it.Next()
	assert.Equal(t, it.Value(), 43)
	assert.Assert(t, !it.Next())
	assert.Assert(t, !it.Next())
}
