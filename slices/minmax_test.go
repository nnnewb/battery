package slices

import (
	"github.com/nnnewb/battery/internal/constraints"
	"reflect"
	"testing"
)

func TestSlice_MinLessFunc(t *testing.T) {
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
	var s = []int{1, 2, 4, 65, 6, 12, 231, 123, 545, 1231, 2345, 5, -100, 1231, 2314}
	tests := []sliceMinLessFuncTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceMinLessFuncArgs{
				less: func(i, j int) bool {
					return i < j
				},
			},
			want:  0,
			want1: false,
		},
		{
			name: "simple",
			s:    s,
			args: sliceMinLessFuncArgs{
				less: func(i, j int) bool {
					return s[i] < s[j]
				},
			},
			want:  -100,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.MinLessFunc(tt.args.less)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinLessFunc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinLessFunc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSlice_MaxLessFunc(t *testing.T) {
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
				less: func(i, j int) bool {
					return i < j
				},
			},
			want:  0,
			want1: false,
		},
		{
			name: "simple",
			s:    s,
			args: sliceMaxLessFuncArgs{
				less: func(i, j int) bool {
					return s[i] < s[j]
				},
			},
			want:  2314,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.MaxLessFunc(tt.args.less)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxLessFunc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MaxLessFunc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestComparableElemSlice_MinLessFunc(t *testing.T) {
	type sliceMinLessFuncArgs struct {
		less func(i, j int) bool
	}
	type sliceMinLessFuncTC[T comparable] struct {
		name  string
		s     ComparableElemSlice[T]
		args  sliceMinLessFuncArgs
		want  T
		want1 bool
	}
	var s = []int{1, 2, 4, 65, 6, 12, 231, 123, 545, 1231, 2345, 5, -100, 1231, 2314}
	tests := []sliceMinLessFuncTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceMinLessFuncArgs{
				less: func(i, j int) bool {
					return i < j
				},
			},
			want:  0,
			want1: false,
		},
		{
			name: "simple",
			s:    s,
			args: sliceMinLessFuncArgs{
				less: func(i, j int) bool {
					return s[i] < s[j]
				},
			},
			want:  -100,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.MinLessFunc(tt.args.less)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinLessFunc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinLessFunc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestComparableElemSlice_MaxLessFunc(t *testing.T) {
	type sliceMaxLessFuncArgs struct {
		less func(i, j int) bool
	}
	type sliceMaxLessFuncTC[T comparable] struct {
		name  string
		s     ComparableElemSlice[T]
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
				less: func(i, j int) bool {
					return i < j
				},
			},
			want:  0,
			want1: false,
		},
		{
			name: "simple",
			s:    s,
			args: sliceMaxLessFuncArgs{
				less: func(i, j int) bool {
					return s[i] < s[j]
				},
			},
			want:  2314,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.MaxLessFunc(tt.args.less)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxLessFunc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MaxLessFunc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderedElemSlice_Min(t *testing.T) {
	type sliceMinArgs struct{}
	type sliceMinTC[T constraints.Ordered] struct {
		name string
		s    OrderedElemSlice[T]
		args sliceMinArgs
		want T
		ok   bool
	}
	tests := []sliceMinTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceMinArgs{},
			want: 0,
			ok:   false,
		},
		{
			name: "simple",
			s:    OrderedElemSlice[int]{1, 2, 4, 65, 6, 12, 231, 123, 545, 1231, 2345, 5, -100, 1231, 2314},
			args: sliceMinArgs{},
			want: -100,
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Min()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Min() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}

func TestOrderedElemSlice_Max(t *testing.T) {
	type sliceMaxArgs struct{}
	type sliceMaxTC[T constraints.Ordered] struct {
		name string
		s    OrderedElemSlice[T]
		args sliceMaxArgs
		want T
		ok   bool
	}
	tests := []sliceMaxTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceMaxArgs{},
			want: 0,
			ok:   false,
		},
		{
			name: "simple",
			s:    []int{1, 2, 4, 65, 6, 12, 231, 123, 545, 1231, 2345, 5, -100, 1231, 2314},
			args: sliceMaxArgs{},
			want: 2345,
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Max()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Max() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}
