package iter

import (
	"fmt"
	assert2 "github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleCollect() {
	numbers := Collect[int](Range[int](0, 3, 1))
	fmt.Println(numbers)
	// Output: [0 1 2]
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

func TestToChannel(t *testing.T) {
	expected := 0
	for number := range ToChannel[int](Lift([]int{1, 2, 3, 4})) {
		expected += 1
		assert2.Equal(t, number, expected)
	}
}

func TestToChannelEmpty(t *testing.T) {
	for range ToChannel[int](Exhausted[int]()) {
		t.Fail()
	}
}
