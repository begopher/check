package valid

type Constraint[T any] interface {
	Evaluate(T) error
}
