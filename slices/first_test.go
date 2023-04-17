package slices

import (
	"github.com/nnnewb/battery/internal/predicate"
	"reflect"
	"testing"
)

func TestSlice_First(t *testing.T) {
	type sliceFirstArgs[T any] struct {
		predicate func(T) bool
	}
	type sliceFirstTC[T any] struct {
		name string
		s    Slice[T]
		args sliceFirstArgs[T]
		want T
		ok   bool
	}
	tests := []sliceFirstTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceFirstArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 0,
			ok:   false,
		}, {
			name: "simple",
			s:    []int{0, 0, 0, 1, 2, 3, 4},
			args: sliceFirstArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 1,
			ok:   true,
		}, {
			name: "simple",
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
			got, got1 := tt.s.First(tt.args.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("First() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}
