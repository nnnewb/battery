package collection

import (
	"fmt"
	"github.com/nnnewb/battery/internal/assert"
	"github.com/nnnewb/battery/iter"
	"strconv"
	"testing"
)

func TestFold(t *testing.T) {
	add := func(a, b int) int { return a + b }
	sum := Fold[int](Iter(iter.Range[int](0, 11, 1)), 0, add)
	assert.Equal(t, sum, 55)

	concat := func(path string, part int) string {
		return path + strconv.Itoa(part) + "/"
	}
	result := Fold[int](Iter(iter.Take[int](iter.Range[int](0, 10, 1), 3)), "/", concat)
	assert.Equal(t, result, "/0/1/2/")
}

func ExampleFold() {
	add := func(a int, b int) int { return a + b }
	sum := Fold[int](Iter(iter.Range[int](0, 4, 1)), 0, add)
	fmt.Println(sum)
	// Output: 6
}
