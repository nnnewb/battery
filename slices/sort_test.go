package slices

import (
	"github.com/nnnewb/battery/internal/constraints"
	"reflect"
	"testing"
)

func TestSlice_SortLessFunc(t *testing.T) {
	type sliceSortLessFuncArgs struct {
		less func(i, j int) bool
	}
	type sliceSortLessFuncTC[T any] struct {
		name string
		s    Slice[T]
		args sliceSortLessFuncArgs
		want Slice[T]
	}
	var s = []int{1, 4, 3, 2, 5, 8, 7, 6, 9, 0}
	tests := []sliceSortLessFuncTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceSortLessFuncArgs{
				less: func(i, j int) bool {
					return s[i] < s[j]
				},
			},
			want: nil,
		},
		{
			name: "simple",
			s:    s,
			args: sliceSortLessFuncArgs{
				less: func(i, j int) bool {
					return s[i] < s[j]
				},
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SortLessFunc(tt.args.less); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortLessFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComparableElemSlice_SortLessFunc(t *testing.T) {
	type sliceSortLessFuncArgs struct {
		less func(i, j int) bool
	}
	type sliceSortLessFuncTC[T comparable] struct {
		name string
		s    ComparableElemSlice[T]
		args sliceSortLessFuncArgs
		want ComparableElemSlice[T]
	}
	var s = []int{1, 4, 3, 2, 5, 8, 7, 6, 9, 0}
	tests := []sliceSortLessFuncTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceSortLessFuncArgs{
				less: func(i, j int) bool {
					return s[i] < s[j]
				},
			},
			want: nil,
		},
		{
			name: "simple",
			s:    s,
			args: sliceSortLessFuncArgs{
				less: func(i, j int) bool {
					return s[i] < s[j]
				},
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SortLessFunc(tt.args.less); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortLessFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedElemSlice_SortLessFunc(t *testing.T) {
	type sliceSortLessFuncTC[T constraints.Ordered] struct {
		name string
		s    OrderedElemSlice[T]
		want OrderedElemSlice[T]
	}
	var s = []int{1, 4, 3, 2, 5, 8, 7, 6, 9, 0}
	tests := []sliceSortLessFuncTC[int]{
		{
			name: "nil",
			s:    nil,
			want: nil,
		},
		{
			name: "simple",
			s:    s,
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortLessFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
