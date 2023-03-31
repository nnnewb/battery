package predicate

import "reflect"

// IsZero 给 Compact 定义的 IsZero 方法，用 reflect 实现。
func IsZero[T any](val T) bool {
	v := reflect.ValueOf(val)
	return !v.IsValid() || v.IsZero()
}
