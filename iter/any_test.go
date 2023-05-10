package iter

import (
	"fmt"
	"testing"

	"github.com/nnnewb/battery/internal/predicate"
)

func ExampleAnyFunc() {
	fmt.Println(AnyFunc(Range(0, 10, 1), predicate.IsZero[int]))
	// output:
	// true
}

func TestAnyFunc(t *testing.T) {
	type anyArgs[T any] struct {
		iterable  Iterator[int]
		predicate func(T) bool
	}
	type anyTestCase[T any] struct {
		name string
		args anyArgs[T]
		want bool
	}
	tests := []anyTestCase[int]{
		{
			name: "expect true",
			args: anyArgs[int]{
				iterable:  Range(-10, 2, 1),
				predicate: predicate.IsPositive[int],
			},
			want: true,
		},
		{
			name: "expect false",
			args: anyArgs[int]{
				iterable:  Range(-10, 1, 1),
				predicate: predicate.IsPositive[int],
			},
			want: false,
		},
		{
			name: "exhausted",
			args: anyArgs[int]{
				iterable:  Exhausted[int](),
				predicate: predicate.IsPositive[int],
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyFunc(tt.args.iterable, tt.args.predicate); got != tt.want {
				t.Errorf("AnyFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
