//go:build go1.18

package slices

import (
	"testing"

	"github.com/nnnewb/battery/internal/predicate"
)

func TestAny(t *testing.T) {
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
			s:         []int{1, 1, 1, 1, 1},
			predicate: predicate.IsPositive[int],
			want:      true,
		},
		{
			name:      "all zeros",
			s:         []int{0, 0, 0, 0, 0},
			predicate: predicate.IsPositive[int],
			want:      false,
		},
		{
			name:      "mixed ones zeros",
			s:         []int{1, 0, 1, 0, 1, 0},
			predicate: predicate.IsPositive[int],
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.s, tt.predicate); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}
