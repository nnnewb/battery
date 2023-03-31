package iter

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleGroupBy() {
	type record struct {
		recordType int
		name       string
	}

	records := []record{
		{1, "name1"},
		{1, "name2"},
		{2, "name3"},
		{2, "name4"},
	}

	fmt.Println(GroupBy(Lift(records), func(r record) int { return r.recordType }))
	// output:
	// map[1:[{1 name1} {1 name2}] 2:[{2 name3} {2 name4}]]
}

func TestGroupBy(t *testing.T) {
	type args[T any, K comparable] struct {
		it      Iterator[int]
		keyFunc func(T) K
	}
	type testCase[T any, K comparable] struct {
		name string
		args args[T, K]
		want map[K][]T
	}
	tests := []testCase[int, int]{
		{
			name: "exhausted",
			args: args[int, int]{
				it:      Exhausted[int](),
				keyFunc: func(i int) int { return i },
			},
			want: make(map[int][]int),
		},
		{
			name: "simple",
			args: args[int, int]{
				it:      Lift([]int{1, 3, 4, 2, 1, 2, 3, 4, 5, 4, 3, 2, 1}),
				keyFunc: func(i int) int { return i },
			},
			want: map[int][]int{
				1: {1, 1, 1},
				2: {2, 2, 2},
				3: {3, 3, 3},
				4: {4, 4, 4},
				5: {5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.it, tt.args.keyFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
