package iter

import (
	"fmt"
	"github.com/nnnewb/battery/internal/predicate"
	"testing"
)

func ExampleAny() {
	fmt.Println(Any[int](Range[int](0, 10, 1), predicate.IsZero[int]))
	// output:
	// true
}

func TestAny(t *testing.T) {
	type args[T any] struct {
		iterable  Iterator[int]
		predicate func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "expect true",
			args: args[int]{
				iterable:  Range[int](-10, 2, 1),
				predicate: predicate.IsPositive[int],
			},
			want: true,
		},
		{
			name: "expect false",
			args: args[int]{
				iterable:  Range[int](-10, 1, 1),
				predicate: predicate.IsPositive[int],
			},
			want: false,
		},
		{
			name: "exhausted",
			args: args[int]{
				iterable:  Exhausted[int](),
				predicate: predicate.IsPositive[int],
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.iterable, tt.args.predicate); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}
