package set

import (
	"reflect"
	"testing"
)

func TestNewSetFromMapKey(t *testing.T) {
	type newSetFromMapKeyArgs[K comparable, V any] struct {
		m map[K]V
	}
	type newSetFromMapKeyTestCase[K comparable, V any] struct {
		name string
		args newSetFromMapKeyArgs[K, V]
		want Set[K]
	}
	tests := []newSetFromMapKeyTestCase[int, int]{
		{
			name: "from nil map",
			args: newSetFromMapKeyArgs[int, int]{},
			want: make(Set[int]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSetFromMapKey(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSetFromMapKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
