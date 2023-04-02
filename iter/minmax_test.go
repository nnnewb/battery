package iter

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nnnewb/battery/internal/constraints"
)

func ExampleMinFunc() {
	fmt.Println(MinFunc(Lift([]int{6, 23, 4, 1, 2, 54, 1, 0}), func(i int) int { return i }))
	// output:
	// 0 true
}

func ExampleMaxFunc() {
	fmt.Println(MaxFunc(Lift([]int{6, 23, 4, 1, 2, 54, 1, 0}), func(i int) int { return i }))
	// output:
	// 54 true
}

func TestMinFunc(t *testing.T) {
	type minArgs[T any, K constraints.Ordered] struct {
		it      Iterator[int]
		keyFunc func(T) K
	}
	type minTestCase[T any, K constraints.Ordered] struct {
		name   string
		args   minArgs[T, K]
		want   T
		wantOk bool
	}
	tests := []minTestCase[int, int]{
		{
			name: "exhausted",
			args: minArgs[int, int]{
				it:      Exhausted[int](),
				keyFunc: func(i int) int { return i },
			},
			want:   0,
			wantOk: false,
		},
		{
			name: "at first",
			args: minArgs[int, int]{
				it:      Lift([]int{-5, 1, 5, 3, 3, 2, 1, -1, -1, -1, -5}),
				keyFunc: func(i int) int { return i },
			},
			want:   -5,
			wantOk: true,
		},
		{
			name: "at end",
			args: minArgs[int, int]{
				it:      Lift([]int{1, 5, 3, 3, 2, 1, -1, -1, -1, -5}),
				keyFunc: func(i int) int { return i },
			},
			want:   -5,
			wantOk: true,
		},
		{
			name: "at middle",
			args: minArgs[int, int]{
				it:      Lift([]int{1, 5, 3, 3, 2, -5, -1, -1, -1, -3}),
				keyFunc: func(i int) int { return i },
			},
			want:   -5,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := MinFunc(tt.args.it, tt.args.keyFunc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinFunc() got = %v, want %v", got, tt.want)
			}
			if ok != tt.wantOk {
				t.Errorf("MinFunc() ok = %v, want %v", ok, tt.wantOk)
			}
		})
	}
}

func TestMaxFunc(t *testing.T) {
	type maxArgs[T any, K constraints.Ordered] struct {
		it      Iterator[int]
		keyFunc func(T) K
	}
	type maxTestCase[T any, K constraints.Ordered] struct {
		name   string
		args   maxArgs[T, K]
		want   T
		wantOk bool
	}
	tests := []maxTestCase[int, int]{
		{
			name: "exhausted",
			args: maxArgs[int, int]{
				it:      Exhausted[int](),
				keyFunc: func(i int) int { return i },
			},
			want:   0,
			wantOk: false,
		},
		{
			name: "at first",
			args: maxArgs[int, int]{
				it:      Lift([]int{5, 1, 5, 3, 3, 2, 1, -1, -1, -1, -5}),
				keyFunc: func(i int) int { return i },
			},
			want:   5,
			wantOk: true,
		},
		{
			name: "at end",
			args: maxArgs[int, int]{
				it:      Lift([]int{1, 0, 3, 3, 2, 1, -1, -1, -1, 5}),
				keyFunc: func(i int) int { return i },
			},
			want:   5,
			wantOk: true,
		},
		{
			name: "at middle",
			args: maxArgs[int, int]{
				it:      Lift([]int{1, 5, 3, 3, 2, -5, -1, -1, -1, -3}),
				keyFunc: func(i int) int { return i },
			},
			want:   5,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MaxFunc(tt.args.it, tt.args.keyFunc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxFunc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantOk {
				t.Errorf("MaxFunc() got1 = %v, want %v", got1, tt.wantOk)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type minArgs[T constraints.Ordered] struct {
		it Iterator[T]
	}
	type minTestCase[T constraints.Ordered] struct {
		name  string
		args  minArgs[T]
		want  T
		found bool
	}
	tests := []minTestCase[int]{
		{
			name: "exhausted",
			args: minArgs[int]{
				it: Exhausted[int](),
			},
			want:  0,
			found: false,
		},
		{
			name: "at first",
			args: minArgs[int]{
				it: Lift([]int{1, 2, 3, 4, 5}),
			},
			want:  1,
			found: true,
		},
		{
			name: "at middle",
			args: minArgs[int]{
				it: Lift([]int{5, 4, 1, 2, 3}),
			},
			want:  1,
			found: true,
		},
		{
			name: "at end",
			args: minArgs[int]{
				it: Lift([]int{5, 4, 3, 2, 1}),
			},
			want:  1,
			found: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := Min(tt.args.it)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("Min() found = %v, want %v", found, tt.found)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type maxArgs[T constraints.Ordered] struct {
		it Iterator[T]
	}
	type maxTestCase[T constraints.Ordered] struct {
		name  string
		args  maxArgs[T]
		want  T
		found bool
	}
	tests := []maxTestCase[int]{
		{
			name: "exhausted",
			args: maxArgs[int]{
				it: Exhausted[int](),
			},
			want:  0,
			found: false,
		},
		{
			name: "at first",
			args: maxArgs[int]{
				it: Lift([]int{5, 4, 3, 2, 1}),
			},
			want:  5,
			found: true,
		},
		{
			name: "at middle",
			args: maxArgs[int]{
				it: Lift([]int{1, 4, 5, 2, 3}),
			},
			want:  5,
			found: true,
		},
		{
			name: "at end",
			args: maxArgs[int]{
				it: Lift([]int{5, 4, 3, 2, 6}),
			},
			want:  6,
			found: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := Max(tt.args.it)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("Max() found = %v, want %v", found, tt.found)
			}
		})
	}
}
