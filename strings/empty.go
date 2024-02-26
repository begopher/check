package strings

import (
	"github.com/begopher/valid"
)

func Empty() valid.Predicate[string] {
	return empty{}
}

type empty struct{}

func (e empty) Evaluate(value string) bool {
	if value == "" {
		return true
	}
	return false
}
