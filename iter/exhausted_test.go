package iter

import (
	"fmt"
	"testing"

	"github.com/nnnewb/battery/internal/assert"
)

func ExampleExhausted() {
	fmt.Println(!Exhausted[int]().Next())
	// Output: true
}

func TestExhausted(t *testing.T) {
	assert.Assert(t, !Exhausted[int]().Next())
}
