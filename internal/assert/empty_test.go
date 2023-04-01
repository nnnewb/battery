package assert

import "testing"

func TestEmpty(t *testing.T) {
	type emptyArgs[T any] struct {
		t     *testing.T
		items []T
	}
	type emptyTestCase[T any] struct {
		name string
		args emptyArgs[T]
	}
	tests := []emptyTestCase[int]{
		{
			name: "empty slice",
			args: emptyArgs[int]{
				t:     t,
				items: []int{},
			},
		},
		{
			name: "nil",
			args: emptyArgs[int]{
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
