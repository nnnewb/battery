package assert

import "testing"

func TestEmpty(t *testing.T) {
	type args[T any] struct {
		t     *testing.T
		items []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "empty slice",
			args: args[int]{
				t:     t,
				items: []int{},
			},
		},
		{
			name: "nil",
			args: args[int]{
				t:     t,
				items: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Empty(tt.args.t, tt.args.items)
		})
	}
}
