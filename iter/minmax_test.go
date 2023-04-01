package iter

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nnnewb/battery/internal/constraints"
)

func ExampleMin() {
	fmt.Println(Min(Lift([]int{6, 23, 4, 1, 2, 54, 1, 0}), func(i int) int { return i }))
	// output:
	// 0 true
}

func ExampleMax() {
	fmt.Println(Max(Lift([]int{6, 23, 4, 1, 2, 54, 1, 0}), func(i int) int { return i }))
	// output:
	// 54 true
}

func TestMin(t *testing.T) {
	type args[T any, K constraints.Ordered] struct {
		it      Iterator[int]
		keyFunc func(T) K
	}
	type testCase[T any, K constraints.Ordered] struct {
		name   string
		args   args[T, K]
		want   T
		wantOk bool
	}
	tests := []testCase[int, int]{
		{
			name: "exhausted",
			args: args[int, int]{
				it:      Exhausted[int](),
				keyFunc: func(i int) int { return i },
			},
			want:   0,
			wantOk: false,
		},
		{
			name: "at first",
			args: args[int, int]{
				it:      Lift([]int{-5, 1, 5, 3, 3, 2, 1, -1, -1, -1, -5}),
				keyFunc: func(i int) int { return i },
			},
			want:   -5,
			wantOk: true,
		},
		{
			name: "at end",
			args: args[int, int]{
				it:      Lift([]int{1, 5, 3, 3, 2, 1, -1, -1, -1, -5}),
				keyFunc: func(i int) int { return i },
			},
			want:   -5,
			wantOk: true,
		},
		{
			name: "at middle",
			args: args[int, int]{
				it:      Lift([]int{1, 5, 3, 3, 2, -5, -1, -1, -1, -3}),
				keyFunc: func(i int) int { return i },
			},
			want:   -5,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Min(tt.args.it, tt.args.keyFunc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantOk {
				t.Errorf("Min() got1 = %v, want %v", got1, tt.wantOk)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args[T any, K constraints.Ordered] struct {
		it      Iterator[int]
		keyFunc func(T) K
	}
	type testCase[T any, K constraints.Ordered] struct {
		name   string
		args   args[T, K]
		want   T
		wantOk bool
	}
	tests := []testCase[int, int]{
		{
			name: "exhausted",
			args: args[int, int]{
				it:      Exhausted[int](),
				keyFunc: func(i int) int { return i },
			},
			want:   0,
			wantOk: false,
		},
		{
			name: "at first",
			args: args[int, int]{
				it:      Lift([]int{5, 1, 5, 3, 3, 2, 1, -1, -1, -1, -5}),
				keyFunc: func(i int) int { return i },
			},
			want:   5,
			wantOk: true,
		},
		{
			name: "at end",
			args: args[int, int]{
				it:      Lift([]int{1, 0, 3, 3, 2, 1, -1, -1, -1, 5}),
				keyFunc: func(i int) int { return i },
			},
			want:   5,
			wantOk: true,
		},
		{
			name: "at middle",
			args: args[int, int]{
				it:      Lift([]int{1, 5, 3, 3, 2, -5, -1, -1, -1, -3}),
				keyFunc: func(i int) int { return i },
			},
			want:   5,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Max(tt.args.it, tt.args.keyFunc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantOk {
				t.Errorf("Max() got1 = %v, want %v", got1, tt.wantOk)
			}
		})
	}
}
