package assert

import (
	"reflect"
	"testing"
)

// Equal 比较两个值是否一致。若不一致则测试失败。用 reflect.DeepEqual 实现。
func Equal[T any](t *testing.T, expr1, expr2 T) {
	t.Helper()

	if !reflect.DeepEqual(expr1, expr2) {
		t.Fatalf("expect %v equals to %v", expr1, expr2)
	}
}
