package iter

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nnnewb/battery/internal/assert"
	"github.com/nnnewb/battery/internal/predicate"
)

func ExampleFilter() {
	it := Filter(Lift([]int{0, 1, 0, 2}), func(i int) bool { return i == 0 })
	for it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// 0
	// 0
}

func TestFilter(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	it := Filter(Range(0, 10, 1), isEven)
	it.Next()
	assert.Equal(t, it.Value(), 0)
	it.Next()
	assert.Equal(t, it.Value(), 2)
}

func TestFilterEmpty(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	it := Filter(Exhausted[int](), isEven)
	assert.Assert(t, !it.Next())
}

func TestFilterExhausted(t *testing.T) {
	delegate := Exhausted[int]()
	it := Filter(delegate, func(_ int) bool { return true })

	assert.Assert(t, !it.Next())
	assert.Assert(t, !it.Next())
}

func TestFilterExhaustedLater(t *testing.T) {
	delegate := Lift([]int{1})
	it := Filter(delegate, func(_ int) bool { return true })

	it.Next()
	assert.Equal(t, it.Value(), 1)
	assert.Assert(t, !it.Next())
}

func TestFilterTableDriven(t *testing.T) {
	type filterArgs[T any] struct {
		iterator Iterator[int]
		isZero   func(T) bool
	}
	type FilterTestCase[T any] struct {
		name string
		args filterArgs[T]
		want []int
	}
	tests := []FilterTestCase[int]{
		{
			name: "empty",
			args: filterArgs[int]{
				iterator: Exhausted[int](),
				isZero:   predicate.IsZero[int],
			},
			want: make([]int, 0),
		},
		{
			name: "shrink to zero",
			args: filterArgs[int]{
				iterator: Range(-10, 0, 1),
				isZero:   predicate.IsPositive[int],
			},
			want: make([]int, 0),
		},
		{
			name: "no change",
			args: filterArgs[int]{
				iterator: Range(1, 10, 1),
				isZero:   predicate.IsPositive[int],
			},
			want: Collect(Range(1, 10, 1)),
		},
		{
			name: "shrink",
			args: filterArgs[int]{
				iterator: Range(-5, 5, 1),
				isZero:   predicate.IsPositive[int],
			},
			want: Collect(Range(1, 5, 1)),
		},
	}
	for _, tt := range tests {
		got := Filter(tt.args.iterator, tt.args.isZero)
		s := Collect(got)
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Compact() = %v, want %v", s, tt.want)
			}
		})
	}
}
