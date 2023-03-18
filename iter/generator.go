package iter

import "github.com/nnnewb/battery/optional"

func gen[T any](f func(exhausted func()) T) <-chan optional.Optional[T] {
	c := make(chan optional.Optional[T], 0)
	go func() {
		var exhausted bool
		for exhausted {
			c <- optional.Value(f(func() {
				exhausted = true
			}))
		}
	}()
	return c
}
