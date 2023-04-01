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
	type args[T any] struct {
		iterable Iterator[int]
		val      int
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
				iterable: Range(-10, 2, 1),
				val:      1,
			},
			want: true,
		},
		{
			name: "expect false",
			args: args[int]{
				iterable: Range(-10, 1, 1),
				val:      1,
			},
			want: false,
		},
		{
			name: "exhausted",
			args: args[int]{
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
