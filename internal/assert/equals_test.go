package assert

import "testing"

func TestEqual(t *testing.T) {
	type args[T comparable] struct {
		t     *testing.T
		expr1 T
		expr2 T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "equals",
			args: args[int]{
				t:     t,
				expr1: 1,
				expr2: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Equal(tt.args.t, tt.args.expr1, tt.args.expr2)
		})
	}
}

func TestSliceEqual(t *testing.T) {
	type args[T comparable] struct {
		t  *testing.T
		s1 []T
		s2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "both empty",
			args: args[int]{
				t:  t,
				s1: []int{},
				s2: []int{},
			},
		},
		{
			name: "both nil",
			args: args[int]{
				t:  t,
				s1: nil,
				s2: nil,
			},
		},
		{
			name: "both have content",
			args: args[int]{
				t:  t,
				s1: []int{1, 2, 3},
				s2: []int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceEqual(tt.args.t, tt.args.s1, tt.args.s2)
		})
	}
}
