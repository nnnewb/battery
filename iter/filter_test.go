package iter

import (
	"fmt"
	"github.com/nnnewb/battery/internal/assert"
	"github.com/nnnewb/battery/internal/predicate"
	"reflect"
	"testing"
)

func ExampleFilter() {
	it := Filter[int](Lift([]int{0, 1, 0, 2}), func(i int) bool { return i == 0 })
	for it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// 0
	// 0
}

func TestFilter(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	it := Filter[int](Range[int](0, 10, 1), isEven)
	it.Next()
	assert.Equal(t, it.Value(), 0)
	it.Next()
	assert.Equal(t, it.Value(), 2)
}

func TestFilterEmpty(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	it := Filter[int](Exhausted[int](), isEven)
	assert.Assert(t, !it.Next())
}

func TestFilterExhausted(t *testing.T) {
	delegate := Exhausted[int]()
	it := Filter[int](delegate, func(_ int) bool { return true })

	assert.Assert(t, !it.Next())
	assert.Assert(t, !it.Next())
}

func TestFilterExhaustedLater(t *testing.T) {
	delegate := Lift([]int{1})
	it := Filter[int](delegate, func(_ int) bool { return true })

	it.Next()
	assert.Equal(t, it.Value(), 1)
	assert.Assert(t, !it.Next())
}

func TestFilterTableDriven(t *testing.T) {
	type args[T any] struct {
		iterator Iterator[int]
		isZero   func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []int
	}
	tests := []testCase[int]{
		{
			name: "empty",
			args: args[int]{
				iterator: Exhausted[int](),
				isZero:   predicate.IsZero[int],
			},
			want: make([]int, 0),
		},
		{
			name: "shrink to zero",
			args: args[int]{
				iterator: Range(-10, 0, 1),
				isZero:   predicate.IsPositive[int],
			},
			want: make([]int, 0),
		},
		{
			name: "no change",
			args: args[int]{
				iterator: Range(1, 10, 1),
				isZero:   predicate.IsPositive[int],
			},
			want: Collect(Range(1, 10, 1)),
		},
		{
			name: "shrink",
			args: args[int]{
				iterator: Range(-5, 5, 1),
				isZero:   predicate.IsPositive[int],
			},
			want: Collect(Range(1, 5, 1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.iterator, tt.args.isZero); !reflect.DeepEqual(Collect(got), tt.want) {
				t.Errorf("Compact() = %v, want %v", Collect(got), tt.want)
			}
		})
	}
}
