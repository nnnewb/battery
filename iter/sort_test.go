package iter

import (
	"github.com/nnnewb/battery/internal/constraints"
	"reflect"
	"testing"
)

func TestSortInt(t *testing.T) {
	type args[T any, K constraints.Ordered] struct {
		it      Iterator[int]
		keyFunc func(T) K
	}
	type testCase[T any, K constraints.Ordered] struct {
		name string
		args args[T, K]
		want []int
	}
	tests := []testCase[int, int]{
		{
			name: "exhausted",
			args: args[int, int]{
				it:      Exhausted[int](),
				keyFunc: func(i int) int { return i },
			},
			want: []int{},
		},
		{
			name: "simple sort",
			args: args[int, int]{
				it:      Lift([]int{1, 5, 3, 2, 6, 7, 4}),
				keyFunc: func(i int) int { return i },
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.it, tt.args.keyFunc); !reflect.DeepEqual(Collect(got), tt.want) {
				t.Errorf("Sort() = %v, want %v", Collect(got), tt.want)
			}
		})
	}
}

func TestSortStruct(t *testing.T) {
	type Record struct {
		pk   int
		data string
	}
	type args[T any, K constraints.Ordered] struct {
		it      Iterator[Record]
		keyFunc func(T) K
	}
	type testCase[T any, K constraints.Ordered] struct {
		name string
		args args[T, K]
		want []Record
	}
	tests := []testCase[Record, int]{
		{
			name: "simple sort",
			args: args[Record, int]{
				it: Lift([]Record{
					{
						pk:   5,
						data: "5",
					},
					{
						pk:   3,
						data: "3",
					},
					{
						pk:   1,
						data: "1",
					},
					{
						pk:   2,
						data: "2",
					},
					{
						pk:   4,
						data: "4",
					},
				}),
				keyFunc: func(record Record) int { return record.pk },
			},
			want: []Record{
				{
					pk:   1,
					data: "1",
				},
				{
					pk:   2,
					data: "2",
				},
				{
					pk:   3,
					data: "3",
				},
				{
					pk:   4,
					data: "4",
				},
				{
					pk:   5,
					data: "5",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.it, tt.args.keyFunc); !reflect.DeepEqual(Collect(got), tt.want) {
				t.Errorf("Sort() = %v, want %v", Collect(got), tt.want)
			}
		})
	}
}
