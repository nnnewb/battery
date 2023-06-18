package slices

import (
	"reflect"
	"testing"
)

func TestSlice_Concat(t *testing.T) {
	type sliceConcatArgs[T any] struct {
		other [][]int
	}
	type sliceConcatTestCase[T any] struct {
		name string
		s    []T
		args sliceConcatArgs[T]
		want []T
	}
	tests := []sliceConcatTestCase[int]{
		{
			name: "nil slice",
			s:    nil,
			args: sliceConcatArgs[int]{
				other: nil,
			},
			want: []int{},
		},
		{
			name: "nil concat with non-nil",
			s:    nil,
			args: sliceConcatArgs[int]{
				other: [][]int{{1, 1, 1, 1}},
			},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "non-nil concat with nil",
			s:    []int{1, 1, 1, 1},
			args: sliceConcatArgs[int]{
				other: nil,
			},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "add two slice",
			s:    []int{1, 1, 1, 1},
			args: sliceConcatArgs[int]{
				other: [][]int{
					{1, 1, 1, 1},
				},
			},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "add more slice",
			s:    []int{1, 2, 3, 4},
			args: sliceConcatArgs[int]{
				other: [][]int{
					{5, 6, 7, 8},
					{9, 10, 11, 12},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.s, tt.args.other...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
