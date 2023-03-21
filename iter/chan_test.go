package iter

import (
	"fmt"
	assert2 "github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleFromChan() {
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	fmt.Println(Collect[int](FromChan(ch)))

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

	it = it.Next()
	assert2.Equal(t, it.Value(), 1)
	it = it.Next()
	assert2.Equal(t, it.Value(), 2)
	it = it.Next()
	assert2.Equal(t, it.Value(), 3)
	it = it.Next()
	assert2.Assert(t, it.Exhausted())
}

func TestFromChannelEmpty(t *testing.T) {
	ch := make(chan int)
	close(ch)
	assert2.Assert(t, FromChan(ch).Next().Exhausted())
}
