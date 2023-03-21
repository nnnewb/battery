package iter

import (
	"fmt"
	"github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleExhausted() {
	fmt.Println(Exhausted[int]().Next().Exhausted())
	// Output: true
}

func TestExhausted(t *testing.T) {
	assert.Assert(t, Exhausted[int]().Next().Exhausted())
}
