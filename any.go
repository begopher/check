package valid

func Any[T any](rules ...Predicate[T]) Predicate[T] {
	for _, rule := range rules {
		if rule == nil {
			panic("check.Any: nil is not allowed")
		}
	}
	return _any[T]{rules}
}

type _any[T any] struct {
	rules []Predicate[T]
}

func (a _any[T]) Evaluate(value T) bool {
	for _, rule := range a.rules {
		if ok := rule.Evaluate(value); ok {
			return true
		}
	}
	return false
}
