package iter

import (
	"fmt"
	"testing"

	"github.com/nnnewb/battery/internal/assert"
)

func ExampleGenerator() {
	positives := Filter(Generator(func() func() (int, bool) {
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
	fmt.Println(Collect(positives))
	// Output: [4 6 4]
}

func TestGenerator(t *testing.T) {
	it := Generator(func() func() (int, bool) {
		i := -1
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

	it.Next()
	assert.Equal(t, it.Value(), 1)
	it.Next()
	assert.Equal(t, it.Value(), 2)
	assert.Assert(t, !it.Next())
}

func TestGeneratorEmpty(t *testing.T) {
	assert.Assert(t, !Lift([]int{}).Next())
}
