package slices

import (
	"reflect"
	"testing"
)

func TestSortLessFunc(t *testing.T) {
	type sliceSortLessFuncArgs struct {
		less func(i, j int) bool
	}
	type sliceSortLessFuncTC[T any] struct {
		name string
		s    []T
		args sliceSortLessFuncArgs
		want []T
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
			if got := SortLessFunc(tt.s, tt.args.less); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortLessFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
