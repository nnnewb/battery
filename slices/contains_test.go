package slices

import "testing"

func TestSlice_Contains(t *testing.T) {
	type sliceContainsArgs[T any] struct {
		v     T
		equal func(T, T) bool
	}
	type sliceContainsTestCase[T any] struct {
		name string
		s    Slice[T]
		args sliceContainsArgs[T]
		want bool
	}
	tests := []sliceContainsTestCase[int]{
		{
			name: "nil slice",
			s:    nil,
			args: sliceContainsArgs[int]{
				v: 0,
				equal: func(a int, b int) bool {
					return a == b
				},
			},
			want: false,
		},
		{
			name: "expect false",
			s:    []int{1, 1, 1, 1},
			args: sliceContainsArgs[int]{
				v: 2,
				equal: func(a int, b int) bool {
					return a == b
				},
			},
			want: false,
		},
		{
			name: "expect true",
			s:    []int{5, 4, 3, 2, 1},
			args: sliceContainsArgs[int]{
				v: 1,
				equal: func(a int, b int) bool {
					return a == b
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.v, tt.args.equal); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
