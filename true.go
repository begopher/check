package check

func True[T any]() Predicate[T] {
	return _true[T]{}
}

type _true[T any] struct{}

func (_true[T]) Evaluate(T) bool {
	return true
}
