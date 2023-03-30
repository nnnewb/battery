package set

import (
	"github.com/nnnewb/battery/internal/assert"
	"reflect"
	"testing"
)

func TestSet_Add(t *testing.T) {
	type args[T comparable] struct {
		values []T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "add single",
			s:    make(Set[int]),
			args: args[int]{
				values: []int{1},
			},
		},
		{
			name: "add multiple",
			s:    make(Set[int]),
			args: args[int]{
				values: []int{1, 2, 3, 4, 5},
			},
		},
		{
			name: "add empty",
			s:    make(Set[int]),
			args: args[int]{
				values: []int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.values...)
			for _, value := range tt.args.values {
				assert.Assert(t, tt.s.Contains(value))
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args[T comparable] struct {
		value T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "simple",
			s:    NewSetFromSlice([]int{1}),
			args: args[int]{
				value: 1,
			},
		},
		{
			name: "remove from empty set",
			s:    make(Set[int]),
			args: args[int]{
				value: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.value)
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args[T comparable] struct {
		v T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "existed",
			s:    NewSetFromSlice([]int{1, 2, 3}),
			args: args[int]{1},
			want: true,
		},
		{
			name: "not existed",
			s:    NewSetFromSlice([]int{1, 2, 3}),
			args: args[int]{4},
			want: false,
		},
		{
			name: "empty set",
			s:    make(Set[int]),
			args: args[int]{4},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_RelativeComplement(t *testing.T) {
	type args[T comparable] struct {
		other Set[T]
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[int]{
		{
			name: "simple superset",
			s:    NewSetFromSlice([]int{1, 2, 3, 4}),
			args: args[int]{
				other: NewSetFromSlice([]int{1, 2, 3}),
			},
			want: NewSetFromSlice([]int{4}),
		},
		{
			name: "simple subset",
			s:    NewSetFromSlice([]int{1, 2, 3}),
			args: args[int]{
				other: NewSetFromSlice([]int{1, 2, 3, 4}),
			},
			want: NewSetFromSlice([]int{}),
		},
		{
			name: "simple intersection",
			s:    NewSetFromSlice([]int{1, 2, 3}),
			args: args[int]{
				other: NewSetFromSlice([]int{2, 3, 4}),
			},
			want: NewSetFromSlice([]int{1}),
		},
		{
			name: "empty set against superset",
			s:    NewSetFromSlice([]int{}),
			args: args[int]{
				other: NewSetFromSlice([]int{1, 2, 3}),
			},
			want: NewSetFromSlice([]int{}),
		},
		{
			name: "superset against empty set",
			s:    NewSetFromSlice([]int{1, 2, 3}),
			args: args[int]{
				other: NewSetFromSlice([]int{}),
			},
			want: NewSetFromSlice([]int{1, 2, 3}),
		},
		{
			name: "empty set against empty set",
			s:    NewSetFromSlice([]int{}),
			args: args[int]{
				other: NewSetFromSlice([]int{}),
			},
			want: NewSetFromSlice([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RelativeComplement(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RelativeComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		want Set[int]
	}
	tests := []testCase[int]{
		{
			name: "empty set",
			want: Set[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
