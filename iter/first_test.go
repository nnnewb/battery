package iter

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nnnewb/battery/internal/predicate"
)

func ExampleFirst() {
	fmt.Println(First(Range(-10, 10, 1), predicate.IsPositive[int]))
	// output:
	// 1 true
}

func TestFirst(t *testing.T) {
	type firstArgs[T any] struct {
		it        Iterator[int]
		predicate func(T) bool
	}
	type firstTestCase[T any] struct {
		name   string
		args   firstArgs[T]
		want   T
		wantOK bool
	}
	tests := []firstTestCase[int]{
		{
			name: "exhausted",
			args: firstArgs[int]{
				it:        Exhausted[int](),
				predicate: predicate.IsPositive[int],
			},
			want:   0,
			wantOK: false,
		},
		{
			name: "at first",
			args: firstArgs[int]{
				it:        Range(1, 10, 1),
				predicate: predicate.IsPositive[int],
			},
			want:   1,
			wantOK: true,
		},
		{
			name: "at middle",
			args: firstArgs[int]{
				it:        Range(-5, 5, 1),
				predicate: predicate.IsPositive[int],
			},
			want:   1,
			wantOK: true,
		},
		{
			name: "at end",
			args: firstArgs[int]{
				it:        Range(-10, 2, 1),
				predicate: predicate.IsPositive[int],
			},
			want:   1,
			wantOK: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := First(tt.args.it, tt.args.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantOK {
				t.Errorf("First() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}
