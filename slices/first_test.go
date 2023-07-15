//go:build go1.18

package slices

import (
	"reflect"
	"testing"

	"github.com/nnnewb/battery/internal/predicate"
)

func TestSlice_First(t *testing.T) {
	type sliceFirstArgs[T any] struct {
		predicate func(T) bool
	}
	type sliceFirstTC[T any] struct {
		name string
		s    []T
		args sliceFirstArgs[T]
		want T
		ok   bool
	}
	tests := []sliceFirstTC[int]{
		{
			name: "nil slice as parameter",
			s:    nil,
			args: sliceFirstArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 0,
			ok:   false,
		}, {
			name: "found case",
			s:    []int{0, 0, 0, 1, 2, 3, 4},
			args: sliceFirstArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 3,
			ok:   true,
		}, {
			name: "not found case",
			s:    []int{0, -1, 0, -2, 0, -3, 0},
			args: sliceFirstArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 0,
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := First(tt.s, tt.args.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("First() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}
