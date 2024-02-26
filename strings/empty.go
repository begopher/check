package strings

import (
	"github.com/begopher/check"
)

func Empty() check.Predicate[string] {
	return empty{}
}

type empty struct{}

func (e empty) Evaluate(value string) bool {
	if value == "" {
		return true
	}
	return false
}
