package check

func Rule[T any](err error, predicates ...Predicate[T]) Constraint[T] {
	if err == nil {
		panic("check.Rule: err cannot be nil")
	}
	if len(predicates) == 0 {
		panic("check.predicates: cannot be empty")
	}
	for _, predicate := range predicates {
		if predicate == nil {
			panic("check.Rule: nil predicate is not allowed")
		}
	}
	return rule[T]{err, predicates}
}

type rule[T any] struct {
	err       error
	predicates []Predicate[T]
}

func (r rule[T]) Evaluate(value T) error {
	for _, predicate := range r.predicates {
		if ok := predicate.Evaluate(value); !ok {
			return r.err
		}
	}
	return nil
}
