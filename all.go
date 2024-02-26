package check

func All[T any](rules ...Predicate[T]) Predicate[T] {
	for _, rule := range rules {
		if rule == nil {
			panic("check.All: nil is not allowed")
		}
	}
	return all[T]{rules}
}

type all[T any] struct {
	rules []Predicate[T]
}

func (a all[T]) Evaluate(value T) bool {
	for _, rule := range a.rules {
		if ok := rule.Evaluate(value); !ok {
			return false
		}
	}
	return true
}
