package iter

import (
	"fmt"
	assert2 "github.com/nnnewb/battery/internal/assert"
	"testing"
)

func ExampleGenerator() {
	positives := Filter[int](Generator[int](func() func() (int, bool) {
		var i int
		values := []int{-1, 4, 6, 4, -5}
		return func() (int, bool) {
			if i+1 < len(values) {
				i++
				r := values[i]
				return r, true
			}
			return 0, false
		}
	}()), func(i int) bool { return i > -1 })
	fmt.Println(Collect[int](positives))
	// Output: [4 6 4]
}

func TestGenerator(t *testing.T) {
	it := Generator(func() func() (int, bool) {
		var i = -1
		values := []int{1, 2}
		return func() (int, bool) {
			if i+1 < len(values) {
				i++
				r := values[i]
				return r, true
			}
			return 0, false
		}
	}())

	it = it.Next()
	assert2.Equal(t, it.Value(), 1)
	it = it.Next()
	assert2.Equal(t, it.Value(), 2)
	it = it.Next()
	assert2.Assert(t, it.Exhausted())
}

func TestGeneratorEmpty(t *testing.T) {
	assert2.Assert(t, Lift([]int{}).Next().Exhausted())
}
