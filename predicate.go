package check

type Predicate[T any] interface {
	Evaluate(T) bool
}
