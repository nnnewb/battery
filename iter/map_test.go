package iter

import (
	"fmt"
	assert2 "github.com/nnnewb/battery/internal/assert"
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
		Map[int](Count(), double),
		4,
	))
	assert2.SliceEqual(t, items, []int{0, 2, 4, 6})
}

func TestMapEmpty(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := Collect[int](Map[int](Exhausted[int](), double))
	assert2.Empty(t, items)
}

func TestMapExhausted(t *testing.T) {
	delegate := Exhausted[int]()
	it := Map[int](delegate, func(t int) float32 { return float32(t) })

	it = it.Next()
	assert2.Assert(t, it.Exhausted())
	it = it.Next()
	assert2.Assert(t, it.Exhausted())
}
