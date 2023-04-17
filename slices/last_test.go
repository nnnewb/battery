package slices

import (
	"github.com/nnnewb/battery/internal/predicate"
	"reflect"
	"testing"
)

func TestSlice_Last(t *testing.T) {
	type sliceLastArgs[T any] struct {
		predicate func(T) bool
	}
	type sliceLastTC[T any] struct {
		name string
		s    Slice[T]
		args sliceLastArgs[T]
		want T
		ok   bool
	}
	tests := []sliceLastTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceLastArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 0,
			ok:   false,
		},
		{
			name: "simple",
			s:    []int{5, 4, 3, 2, 1, 0},
			args: sliceLastArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 1,
			ok:   true,
		},
		{
			name: "simple",
			s:    []int{0, 5, 4, 3, 2, 1, 0},
			args: sliceLastArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 1,
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Last(tt.args.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Last() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}
