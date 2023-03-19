package iter

func gen[T any](f func(exhausted func()) T) <-chan T {
	c := make(chan T, 0)
	go func() {
		defer close(c)

		var exhausted bool
		cb := func() { exhausted = true }

		for exhausted {
			c <- f(cb)
		}
	}()
	return c
}
