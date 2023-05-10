package iter

import (
	"fmt"
	"testing"
)

func ExampleContainsFunc() {
	equal := func(a int, b int) bool { return a == b }
	fmt.Println(ContainsFunc(Range(0, 10, 1), 0, equal))
	// output:
	// true
}

func TestContainsFunc(t *testing.T) {
	equal := func(a int, b int) bool { return a == b }
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
			if got := ContainsFunc(tt.args.iterable, tt.args.val, equal); got != tt.want {
				t.Errorf("AnyFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
