//go:build go1.18

package slices

import (
	"math/rand"
	"testing"

	"github.com/nnnewb/battery/internal/predicate"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name      string
		s         []int
		predicate func(int) bool
		want      bool
	}{
		{
			name:      "nil slice",
			s:         nil,
			predicate: predicate.IsPositive[int],
			want:      false,
		},
		{
			name:      "all ones",
			s:         []int{1, 1, 1, 1},
			predicate: predicate.IsPositive[int],
			want:      true,
		},
		{
			name:      "all zeros",
			s:         []int{0, 0, 0, 0},
			predicate: predicate.IsPositive[int],
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.s, tt.predicate); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzAll(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint16) {
		s := make([]int, size)
		for i := 0; i < len(s); i++ {
			s[i] = rand.Int()
		}
		All(s, predicate.IsPositive[int])
	})
}
