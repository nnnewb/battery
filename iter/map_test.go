package iter

import (
	"fmt"
	"github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleMap() {
	double := func(a int) int { return a * 2 }
	items := Collect[int](Map[int](Lift([]int{0, 1, 2, 3}), double))

	fmt.Println(items)
	// Output: [0 2 4 6]
}

func TestMap(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := Collect[int](Take[int](
		Map[int](Range[int](0, 10, 1), double),
		4,
	))
	assert.SliceEqual(t, items, []int{0, 2, 4, 6})
}

func TestMapEmpty(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := Collect[int](Map[int](Exhausted[int](), double))
	assert.Empty(t, items)
}

func TestMapExhausted(t *testing.T) {
	delegate := Exhausted[int]()
	it := Map[int](delegate, func(t int) float32 { return float32(t) })

	assert.Assert(t, !it.Next())
	assert.Assert(t, !it.Next())
}
