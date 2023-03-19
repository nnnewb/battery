package iter

import (
	"fmt"
	"github.com/nnnewb/battery/assert"
	"testing"
)

func ExampleSkip() {
	it := Skip[int](Count(), 2)
	it = it.Next()
	fmt.Println(it.Value())
	// Output: 3
}

func TestSkip(t *testing.T) {
	counter := Skip[int](Count(), 2)
	assert.Equal(t, counter.Next().Value(), 3)
}

func TestSkipExhausted(t *testing.T) {
	delegate := Mock[int]()
	it := Skip[int](delegate, 5)

	it = it.Next()
	assert.Assert(t, it.Exhausted())
	it = it.Next()
	assert.Assert(t, it.Exhausted())

	assert.Equal(t, it.(skipIter[int]).underlying.(MockIterator[int]).nextCallCount, 6)
}

func TestSkipExhaustedLater(t *testing.T) {
	delegate := Mock[int](42, 43)
	it := Skip[int](delegate, 1)

	it = it.Next()
	assert.Equal(t, it.Value(), 43)
	it = it.Next()
	assert.Assert(t, it.Exhausted())
	it = it.Next()
	assert.Assert(t, it.Exhausted())
	assert.Equal(t, it.(skipIter[int]).underlying.(MockIterator[int]).nextCallCount, 3)
}
