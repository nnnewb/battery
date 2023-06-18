package slices

import "testing"

func TestContains(t *testing.T) {
	// Positive test case
	s := []string{"apple", "banana", "cherry"}
	v := "banana"
	expected := true
	actual := Contains(s, v)
	if actual != expected {
		t.Errorf("Contains(%v, %s) = %t; expected %t", s, v, actual, expected)
	}
	// Negative test case
	s = []string{"apple", "banana", "cherry"}
	v = "orange"
	expected = false
	actual = Contains(s, v)
	if actual != expected {
		t.Errorf("Contains(%v, %s) = %t; expected %t", s, v, actual, expected)
	}
}

func TestContainsFunc(t *testing.T) {
	tests := []struct {
		name  string
		s     []int
		v     int
		equal func(int, int) bool
		want  bool
	}{
		{
			name: "nil slice",
			s:    nil,
			v:    0,
			equal: func(a int, b int) bool {
				return a == b
			},
			want: false,
		},
		{
			name: "expect false",
			s:    []int{1, 1, 1, 1},
			v:    2,
			equal: func(a int, b int) bool {
				return a == b
			},
			want: false,
		},
		{
			name: "expect true",
			s:    []int{5, 4, 3, 2, 1},
			v:    1,
			equal: func(a int, b int) bool {
				return a == b
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsFunc(tt.s, tt.v, tt.equal); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
