package iter

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/nnnewb/battery/internal/assert"
)

func TestReduce(t *testing.T) {
	add := func(a, b int) int { return a + b }
	sum := Reduce(Range(0, 11, 1), 0, add)
	assert.Equal(t, sum, 55)

	concat := func(path string, part int) string {
		return path + strconv.Itoa(part) + "/"
	}
	result := Reduce(Take(Range(0, 10, 1), 3), "/", concat)
	assert.Equal(t, result, "/0/1/2/")
}

func ExampleReduce() {
	add := func(a int, b int) int { return a + b }
	sum := Reduce(Range(0, 4, 1), 0, add)
	fmt.Println(sum)
	// Output: 6
}
