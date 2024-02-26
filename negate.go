package check

func Negate[T any](rule Predicate[T]) Predicate[T] {
	if rule == nil {
		panic("check.Negate: nil is not allowed")
	}
	return negate[T]{rule}
}

type negate[T any] struct {
	rule Predicate[T]
}

func (n negate[T]) Evaluate(value T) bool {
	return !n.rule.Evaluate(value)
}
