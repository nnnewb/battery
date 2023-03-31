package predicate

import "github.com/nnnewb/battery/internal/constraints"

func IsPositive[T constraints.Numeric](i T) bool {
	return i > 0
}

func IsNegative[T constraints.Numeric](i T) bool {
	return i < 0
}
