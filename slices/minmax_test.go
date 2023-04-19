package slices

import (
	"github.com/nnnewb/battery/internal/constraints"
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	type minArgs[T constraints.Ordered] struct {
		s Slice[T]
	}
	type minTC[T constraints.Ordered] struct {
		name string
		args minArgs[T]
		want T
		ok   bool
	}
	tests := []minTC[int]{
		{
			name: "nil",
			args: minArgs[int]{
				s: nil,
			},
			want: 0,
			ok:   false,
		},
		{
			name: "simple",
			args: minArgs[int]{
				s: []int{1, 2, 3, 3, 2, 1, 6, 7, 23, 4, 5, 7, -10, 8, 23, 4, 6, 87, 45, 23, 0},
			},
			want: -10,
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Min(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Min() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type maxArgs[T constraints.Ordered] struct {
		s Slice[T]
	}
	type maxTC[T constraints.Ordered] struct {
		name string
		args maxArgs[T]
		want T
		ok   bool
	}
	tests := []maxTC[int]{
		{
			name: "nil",
			args: maxArgs[int]{
				s: nil,
			},
			want: 0,
			ok:   false,
		},
		{
			name: "simple",
			args: maxArgs[int]{
				s: []int{1, 5, 1, 23, 5, 77, -10, 345, 2, 43, 4, 24, 31, 34, 32145, 234, 53, 245},
			},
			want: 32145,
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Max(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Max() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}

func TestMinKeyFunc(t *testing.T) {
	type minKeyFuncArgs[T any, K constraints.Ordered] struct {
		s       Slice[T]
		keyFunc func(T) K
	}
	type minKeyFuncTC[T any, K constraints.Ordered] struct {
		name  string
		args  minKeyFuncArgs[T, K]
		want  T
		want1 bool
	}
	tests := []minKeyFuncTC[int, int]{
		{
			name: "nil",
			args: minKeyFuncArgs[int, int]{
				s: nil,
				keyFunc: func(i int) int {
					return i
				},
			},
			want:  0,
			want1: false,
		},
		{
			name: "simple",
			args: minKeyFuncArgs[int, int]{
				s: []int{1, 2, 3, 4, 3, 2, 1, 0, -11, 12, 3, 444, 22, 33, 4, 5, 66},
				keyFunc: func(i int) int {
					return i
				},
			},
			want:  -11,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MinKeyFunc(tt.args.s, tt.args.keyFunc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinKeyFunc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinKeyFunc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMaxKeyFunc(t *testing.T) {
	type maxKeyFuncArgs[T any, K constraints.Ordered] struct {
		s       Slice[T]
		keyFunc func(T) K
	}
	type maxKeyFuncTC[T any, K constraints.Ordered] struct {
		name  string
		args  maxKeyFuncArgs[T, K]
		want  T
		want1 bool
	}
	tests := []maxKeyFuncTC[int, int]{
		{
			name: "nil",
			args: maxKeyFuncArgs[int, int]{
				s: nil,
				keyFunc: func(i int) int {
					return i
				},
			},
			want:  0,
			want1: false,
		},
		{
			name: "simple",
			args: maxKeyFuncArgs[int, int]{
				s: []int{1, 2, 3, 4, 3, 2, 1, 0, -11, 12, 3, 444, 22, 33, 4, 5, 66},
				keyFunc: func(i int) int {
					return i
				},
			},
			want:  444,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MaxKeyFunc(tt.args.s, tt.args.keyFunc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxKeyFunc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MaxKeyFunc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
