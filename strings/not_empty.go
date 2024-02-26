package strings

import (
	"github.com/begopher/valid"
)

func NotEmpty() valid.Predicate[string] {
	return notEmpty{}
}

type notEmpty struct{}

func (notEmpty) Evaluate(value string) bool {
	if value != "" {
		return true
	}
	return false
}
