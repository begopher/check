package valid

type Predicate[T any] interface {
	Evaluate(T) bool
}
