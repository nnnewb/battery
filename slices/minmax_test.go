package slices

import (
	"reflect"
	"testing"
)

func TestSlice_Min(t *testing.T) {
	type sliceMinLessFuncArgs struct {
		less func(i, j int) bool
	}
	type sliceMinLessFuncTC[T any] struct {
		name  string
		s     Slice[T]
		args  sliceMinLessFuncArgs
		want  T
		want1 bool
	}
	tests := []sliceMinLessFuncTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceMinLessFuncArgs{
				less: func(i, j int) bool { return i < j },
			},
			want:  0,
			want1: false,
		},
		{
			name: "simple",
			s:    []int{1, 2, 4, 65, 6, 12, 231, 123, 545, 1231, 2345, 5, -100, 1231, 2314},
			args: sliceMinLessFuncArgs{
				less: func(i, j int) bool { return i < j },
			},
			want:  -100,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Min(tt.args.less)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinLessFunc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinLessFunc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSlice_Max(t *testing.T) {
	type sliceMaxLessFuncArgs struct {
		less func(i, j int) bool
	}
	type sliceMaxLessFuncTC[T any] struct {
		name  string
		s     Slice[T]
		args  sliceMaxLessFuncArgs
		want  T
		want1 bool
	}
	var s = []int{1, 2, 4, 65, 6, 12, 231, 123, 545, 1231, 2345, 5, -100, 1231, 2314}
	tests := []sliceMaxLessFuncTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceMaxLessFuncArgs{
				less: func(i, j int) bool { return i < j },
			},
			want:  0,
			want1: false,
		},
		{
			name: "simple",
			s:    s,
			args: sliceMaxLessFuncArgs{
				less: func(i, j int) bool { return i < j },
			},
			want:  2345,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Max(tt.args.less)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxLessFunc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MaxLessFunc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
