package iter

import (
	"fmt"
	"github.com/nnnewb/battery/internal/assert"
	"testing"
)

type tuple[T1, T2 any] struct {
	One T1
	Two T2
}

func makeTuple[T1, T2 any](a T1, b T2) tuple[T1, T2] {
	return tuple[T1, T2]{
		One: a,
		Two: b,
	}
}

func ExampleZip() {
	isEven := func(a int) bool { return a%2 == 0 }
	isOdd := func(a int) bool { return !isEven(a) }
	evens := Filter[int](Range[int](0, 10, 1), isEven)
	odds := Filter[int](Range[int](0, 10, 1), isOdd)

	zipped := Collect[tuple[int, int]](
		Take[tuple[int, int]](Zip[int, int](evens, odds, makeTuple[int, int]), 3),
	)
	fmt.Println(zipped)
	// Output: [{0 1} {2 3} {4 5}]
}

func TestZip(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	isOdd := func(a int) bool { return !isEven(a) }
	evens := Filter[int](Range[int](0, 10, 1), isEven)
	odds := Filter[int](Range[int](0, 10, 1), isOdd)

	zipped := Collect[tuple[int, int]](
		Take[tuple[int, int]](Zip[int, int](evens, odds, makeTuple[int, int]), 3),
	)

	assert.SliceEqual(t, zipped, []tuple[int, int]{{0, 1}, {2, 3}, {4, 5}})
}
