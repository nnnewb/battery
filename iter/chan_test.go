package iter

import (
	"fmt"
	"testing"

	"github.com/nnnewb/battery/internal/assert"
)

func ExampleFromChan() {
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	fmt.Println(Collect(FromChan(ch)))

	// Output: [1 2]
}

func TestFromChannel(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	it := FromChan(ch)

	it.Next()
	assert.Equal(t, it.Value(), 1)
	it.Next()
	assert.Equal(t, it.Value(), 2)
	it.Next()
	assert.Equal(t, it.Value(), 3)
	assert.Assert(t, !it.Next())
}

func TestFromChannelEmpty(t *testing.T) {
	ch := make(chan int)
	close(ch)
	assert.Assert(t, !FromChan(ch).Next())
}
