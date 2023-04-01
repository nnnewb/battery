package iter

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleCountBy() {
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

	fmt.Println(CountBy(Lift(records), func(r record) int { return r.recordType }))
	// output:
	// map[1:2 2:2]
}

func TestCountBy(t *testing.T) {
	type countByArgs[T any, K comparable] struct {
		it      Iterator[int]
		keyFunc func(T) K
	}
	type countByTestCase[T any, K comparable] struct {
		name string
		args countByArgs[T, K]
		want map[K]int
	}
	tests := []countByTestCase[int, int]{
		{
			name: "exhausted",
			args: countByArgs[int, int]{
				it:      Exhausted[int](),
				keyFunc: func(i int) int { return i },
			},
			want: make(map[int]int),
		},
		{
			name: "simple",
			args: countByArgs[int, int]{
				it:      Lift([]int{4, 4, 3, 2, 1, 1, 2, 3, 5}),
				keyFunc: func(i int) int { return i },
			},
			want: map[int]int{
				4: 2,
				3: 2,
				2: 2,
				1: 2,
				5: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBy(tt.args.it, tt.args.keyFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
