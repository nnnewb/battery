package iter

import (
	"fmt"
	"testing"
)

func ExampleContains() {
	fmt.Println(Contains(Range(0, 10, 1), 0))
	// output:
	// true
}

func TestContains(t *testing.T) {
	type containsArgs[T any] struct {
		iterable Iterator[int]
		val      int
	}
	type containsTestCase[T any] struct {
		name string
		args containsArgs[T]
		want bool
	}
	tests := []containsTestCase[int]{
		{
			name: "expect true",
			args: containsArgs[int]{
				iterable: Range(-10, 2, 1),
				val:      1,
			},
			want: true,
		},
		{
			name: "expect false",
			args: containsArgs[int]{
				iterable: Range(-10, 1, 1),
				val:      1,
			},
			want: false,
		},
		{
			name: "exhausted",
			args: containsArgs[int]{
				iterable: Exhausted[int](),
				val:      1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.iterable, tt.args.val); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}
