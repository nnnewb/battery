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
	sum := Fold[int](Iter(iter.Take[int](iter.Count(), 11)), 0, add)
	assert.Equal(t, sum, 55)

	concat := func(path string, part int) string {
		return path + strconv.Itoa(part) + "/"
	}
	result := Fold[int](Iter(iter.Take[int](iter.Count(), 3)), "/", concat)
	assert.Equal(t, result, "/0/1/2/")
}

func ExampleFold() {
	add := func(a int, b int) int { return a + b }
	sum := Fold[int](Iter(iter.Take[int](iter.Count(), 4)), 0, add)
	fmt.Println(sum)
	// Output: 6
}
