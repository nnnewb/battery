package iter

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"

	assert2 "github.com/nnnewb/battery/internal/assert"
)

func ExampleLift() {
	positives := Filter(Lift([]int{-1, 4, 6, 4, -5}), func(i int) bool { return i > -1 })
	fmt.Println(Collect(positives))
	// Output: [4 6 4]
}

func TestLift(t *testing.T) {
	it := Lift([]int{1, 2})
	it.Next()
	assert2.Equal(t, it.Value(), 1)
	it.Next()
	assert2.Equal(t, it.Value(), 2)
	assert2.Assert(t, !it.Next())
}

func TestLiftEmpty(t *testing.T) {
	assert2.Assert(t, !Lift([]int{}).Next())
}

// BenchmarkLiftIteration 验证下 Lift 等 Generator 形式创建的迭代器对比下标访问的性能损失有多大
//
//	cpu: AMD Ryzen 7 4800H with Radeon Graphics
//	BenchmarkLiftIteration/iteration_by_lift-16               1000000000 0.3121 ns/op  0 B/op 0 allocs/op
//	BenchmarkLiftIteration/iteration_by_ordinary_for-range-16 1000000000 0.02552 ns/op 0 B/op 0 allocs/op
func BenchmarkLiftIteration(b *testing.B) {
	data := make([]byte, 1024*1024*100)
	n, err := io.ReadAtLeast(rand.Reader, data, len(data))
	if err != nil {
		b.Fatal(err)
	}
	if n < len(data) {
		b.Fatal("can not read enough random bytes")
	}

	b.ResetTimer()
	b.Run("iteration by lift", func(b *testing.B) {
		it := Lift(data)
		for it.Next() {
		}
	})
	b.Run("iteration by ordinary for loop", func(b *testing.B) {
		for i := 0; i < len(data); i++ {
		}
	})
}
