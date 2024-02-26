package check

func False[T any]() Predicate[T] {
	return _false[T]{}
}

type _false[T any] struct{}

func (_false[T]) Evaluate(T) bool {
	return false
}
