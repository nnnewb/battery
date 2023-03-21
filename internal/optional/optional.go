package optional

type Optional[T any] interface {
	HasValue() bool
	Value() T
}

type Option[T any] struct {
	hasValue bool
	v        T
}

func Value[T any](v T) Option[T] {
	return Option[T]{
		hasValue: true,
		v:        v,
	}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) HasValue() bool {
	return o.hasValue
}

func (o Option[T]) Value() T {
	return o.v
}
