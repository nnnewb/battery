package iter

import (
	"fmt"
	"testing"
)

func ExampleAllFunc() {
	positive := func(i int) bool { return i > 0 }
	fmt.Println(AllFunc(Range(1, 10, 1), positive))
	// output:
	// true
}

func TestAllFunc(t *testing.T) {
	type allArgs[T any] struct {
		iterable  Iterator[int]
		predicate func(T) bool
	}
	type allTestCase[T any] struct {
		name string
		args allArgs[T]
		want bool
	}
	tests := []allTestCase[int]{
		{
			name: "empty",
			args: allArgs[int]{
				iterable:  Exhausted[int](),
				predicate: func(i int) bool { return i > 0 },
			},
			want: true,
		},
		{
			name: "expect true",
			args: allArgs[int]{
				iterable:  Range(1, 5, 1),
				predicate: func(i int) bool { return i > 0 },
			},
			want: true,
		},
		{
			name: "expect false",
			args: allArgs[int]{
				iterable:  Range(-5, 0, 1),
				predicate: func(i int) bool { return i > 0 },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllFunc(tt.args.iterable, tt.args.predicate); got != tt.want {
				t.Errorf("AllFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
