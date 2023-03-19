package iter

import (
	"fmt"
	"github.com/nnnewb/battery/assert"
	"testing"
)

func ExampleCount() {
	it := Count()

	it = it.Next()
	fmt.Println(it.Value())
	it = it.Next()
	fmt.Println(it.Value())
	it = it.Next()
	fmt.Println(it.Value())

	// Output:
	// 0
	// 1
	// 2
}

func TestCount(t *testing.T) {
	it := Count()

	it = it.Next()
	assert.Equal(t, it.Value(), 0)
	it = it.Next()
	assert.Equal(t, it.Value(), 1)
	it = it.Next()
	assert.Equal(t, it.Value(), 2)
}
