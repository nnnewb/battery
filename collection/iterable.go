package collection

import "github.com/nnnewb/battery/iter"

type Iterable[T any] interface {
	Iterator() iter.Iterator[T]
}
