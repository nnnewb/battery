package slices

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	type mapArgs[T any, R any] struct {
		s []T
		f func(T) R
	}
	type mapTC[T any, R any] struct {
		name string
		args mapArgs[T, R]
		want []R
	}
	tests := []mapTC[int, int]{
		{
			name: "nil",
			args: mapArgs[int, int]{
				s: nil,
				f: func(i int) int {
					return i + 1
				},
			},
			want: []int{},
		},
		{
			name: "add one",
			args: mapArgs[int, int]{
				s: []int{1, 2, 3, 4, 5},
				f: func(i int) int {
					return i + 1
				},
			},
			want: []int{2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapTypeConv(t *testing.T) {
	type mapArgs[T any, R any] struct {
		s []T
		f func(T) R
	}
	type mapTC[T any, R any] struct {
		name string
		args mapArgs[T, R]
		want []R
	}
	tests := []mapTC[int, string]{
		{
			name: "nil",
			args: mapArgs[int, string]{
				s: nil,
				f: strconv.Itoa,
			},
			want: []string{},
		},
		{
			name: "itoa",
			args: mapArgs[int, string]{
				s: []int{1, 2, 3, 4, 5},
				f: strconv.Itoa,
			},
			want: []string{"1", "2", "3", "4", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
