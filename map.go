package valid

func Map[F, T any](mapping func(F) T, rule Predicate[T]) Predicate[F] {
	if rule == nil {
		panic("check.Map: rule cannot be nil")
	}
	if mapping == nil {
		panic("check.Map: mapping cannot be nil")
	}
	return _map[F, T]{rule, mapping}
}

type _map[F, T any] struct {
	rule    Predicate[T]
	mapping func(F) T
}

func (m _map[F, T]) Evaluate(from F) bool {
	to := m.mapping(from)
	return m.rule.Evaluate(to)
}
