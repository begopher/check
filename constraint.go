package check

type Constraint[T any] interface {
	Evaluate(T) error
}
