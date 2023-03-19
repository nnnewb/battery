package iter

import (
	"fmt"
	"github.com/nnnewb/battery/assert"
	"strconv"
	"testing"
)

func ExampleCollect() {
	numbers := Collect[int](Take[int](Count(), 3))
	fmt.Println(numbers)
	// Output: [0 1 2]
}

func ExampleFold() {
	sum := Fold[int](Take[int](Count(), 4), 0, func(a int, b int) int { return a + b })
	fmt.Println(sum)
	// Output: 6
}

func ExampleToChannel() {
	for number := range ToChannel[int](Lift([]int{1, 2, 3})) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestCollect(t *testing.T) {
	items := Collect[int](Take[int](Count(), 5))
	assert.SliceEqual(t, items, []int{0, 1, 2, 3, 4})
}

func TestCollectEmpty(t *testing.T) {
	items := Collect[int](Take[int](Count(), 0))
	assert.Empty(t, items)
}

func TestFold(t *testing.T) {
	add := func(a, b int) int { return a + b }
	sum := Fold[int](Take[int](Count(), 11), 0, add)
	assert.Equal(t, sum, 55)

	concat := func(path string, part int) string {
		return path + strconv.Itoa(part) + "/"
	}
	result := Fold[int](Take[int](Count(), 3), "/", concat)
	assert.Equal(t, result, "/0/1/2/")
}

func TestToChannel(t *testing.T) {
	expected := 0
	for number := range ToChannel[int](Lift([]int{1, 2, 3, 4})) {
		expected += 1
		assert.Equal(t, number, expected)
	}
}

func TestToChannelEmpty(t *testing.T) {
	for range ToChannel[int](Exhausted[int]()) {
		t.Fail()
	}
}
