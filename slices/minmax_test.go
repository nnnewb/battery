package slices

import (
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	type sliceMinLessFuncTC[T any] struct {
		name    string
		s       []T
		wantIdx int
		wantVal T
		wantOK  bool
	}
	tests := []sliceMinLessFuncTC[int]{
		{
			name:    "nil",
			s:       nil,
			wantIdx: 0,
			wantVal: 0,
			wantOK:  false,
		},
		{
			name:    "simple",
			s:       []int{1, 2, 4, 65, 6, 12, 231, 123, 545, 1231, 2345, 5, -100, 1231, 2314},
			wantIdx: 12,
			wantVal: -100,
			wantOK:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdx, gotVal, gotOK := Min(tt.s)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("MinLessFunc() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if gotIdx != tt.wantIdx {
				t.Errorf("MinLessFunc() gotIdx = %v, want %v", gotIdx, tt.wantIdx)
			}
			if gotOK != tt.wantOK {
				t.Errorf("MinLessFunc() gotOK = %v, want %v", gotOK, tt.wantOK)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type sliceMaxLessFuncTC[T any] struct {
		name    string
		s       []T
		wantIdx int
		wantVal T
		wantOK  bool
	}
	tests := []sliceMaxLessFuncTC[int]{
		{
			name:    "nil",
			s:       nil,
			wantIdx: 0,
			wantVal: 0,
			wantOK:  false,
		},
		{
			name:    "simple",
			s:       []int{1, 2, 4, 65, 6, 12, 231, 123, 545, 1231, 2345, 5, -100, 1231, 2314},
			wantIdx: 10,
			wantVal: 2345,
			wantOK:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdx, gotVal, gotOK := Max(tt.s)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("MaxLessFunc() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if gotOK != tt.wantOK {
				t.Errorf("MaxLessFunc() gotOK = %v, want %v", gotOK, tt.wantOK)
			}
			if gotIdx != tt.wantIdx {
				t.Errorf("MaxLessFunc() gotIdx = %v, want %v", gotIdx, tt.wantIdx)
			}
		})
	}
}
