package iter

import (
	"fmt"
	"testing"

	assert "github.com/nnnewb/battery/internal/assert"
)

func ExampleTake() {
	it := Take(Range(1, 10, 1), 2)
	it.Next()
	fmt.Println(it.Value())
	it.Next()
	fmt.Println(it.Value())
	fmt.Println(!it.Next())
	// Output:
	// 1
	// 2
	// true
}

func TestTakeIter(t *testing.T) {
	it := Take(Range(0, 10, 1), 2)

	it.Next()
	assert.Equal(t, it.Value(), 0)
	it.Next()
	assert.Equal(t, it.Value(), 1)
	assert.Assert(t, !it.Next())
}

func TestTakeIterEmpty(t *testing.T) {
	it := Take(Range(0, 10, 1), 0)
	assert.Assert(t, !it.Next())
}

func TestTakeExhausted(t *testing.T) {
	delegate := Exhausted[int]()
	it := Take(delegate, 10)

	assert.Assert(t, !it.Next())
	assert.Assert(t, !it.Next())
}
