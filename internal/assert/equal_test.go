package assert

import "testing"

func TestEqual(t *testing.T) {
	type equalArgs[T comparable] struct {
		t     *testing.T
		expr1 T
		expr2 T
	}
	type equalTestCase[T comparable] struct {
		name string
		args equalArgs[T]
	}
	tests := []equalTestCase[int]{
		{
			name: "equals",
			args: equalArgs[int]{
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
