package iter

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nnnewb/battery/internal/predicate"
)

func ExampleLast() {
	fmt.Println(Last(Lift([]int{1, 2, 3, 4, 5}), predicate.IsPositive[int]))
	// output:
	// 5 true
}

func TestLast(t *testing.T) {
	type args[T any] struct {
		it        Iterator[int]
		predicate func(T) bool
	}
	type testCase[T any] struct {
		name   string
		args   args[T]
		want   T
		wantOk bool
	}
	tests := []testCase[int]{
		{
			name: "exhausted",
			args: args[int]{
				it:        Exhausted[int](),
				predicate: predicate.IsPositive[int],
			},
			want:   0,
			wantOk: false,
		},
		{
			name: "at end",
			args: args[int]{
				it:        Range(1, 10, 1),
				predicate: predicate.IsPositive[int],
			},
			want:   9,
			wantOk: true,
		},
		{
			name: "at middle",
			args: args[int]{
				it:        Lift([]int{-1, 0, 1, -2, -3}),
				predicate: predicate.IsPositive[int],
			},
			want:   1,
			wantOk: true,
		},
		{
			name: "at first",
			args: args[int]{
				it:        Lift([]int{1, 0, -1, -2, -3}),
				predicate: predicate.IsPositive[int],
			},
			want:   1,
			wantOk: true,
		},
		{
			name: "find zero",
			args: args[int]{
				it:        Range(0, 1, 1),
				predicate: predicate.IsZero[int],
			},
			want:   0,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Last(tt.args.it, tt.args.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantOk {
				t.Errorf("Last() got1 = %v, want %v", got1, tt.wantOk)
			}
		})
	}
}
