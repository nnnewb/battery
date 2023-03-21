package iter

import (
	"fmt"
	assert2 "github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleSkip() {
	it := Skip[int](Count(), 2)
	it = it.Next()
	fmt.Println(it.Value())
	// Output: 2
}

func TestSkip(t *testing.T) {
	counter := Skip[int](Count(), 2)
	assert2.Equal(t, counter.Next().Value(), 2)
}

func TestSkipExhausted(t *testing.T) {
	delegate := Exhausted[int]()
	it := Skip[int](delegate, 5)

	it = it.Next()
	assert2.Assert(t, it.Exhausted())
	it = it.Next()
	assert2.Assert(t, it.Exhausted())
}

func TestSkipExhaustedLater(t *testing.T) {
	delegate := Lift([]int{42, 43})
	it := Skip[int](delegate, 1)

	it = it.Next()
	assert2.Equal(t, it.Value(), 43)
	it = it.Next()
	assert2.Assert(t, it.Exhausted())
	it = it.Next()
	assert2.Assert(t, it.Exhausted())
}
