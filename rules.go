package valid

func Rules[T any](many ...Constraint[T]) Constraint[T] {
	for _, c := range many {
		if c == nil {
			panic("check.Rules: nil constraint is not allowed")
		}
	}
	return rules[T]{many}
}

type rules[T any] struct {
	many []Constraint[T]
}

func (r rules[T]) Evaluate(value T) error {
	for _, constraint := range r.many {
		if err := constraint.Evaluate(value); err != nil {
			return err
		}
	}
	return nil
}
