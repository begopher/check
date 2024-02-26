package strings

import (
	"github.com/begopher/valid"
)

func Equal(many ...string) valid.Predicate[string] {
	if len(many) == 0 {
		panic("strings.Is: many cannot be empty")
	}
	return equal{many}
}

type equal struct {
	many []string
}

func (e equal) Evaluate(value string) bool {
	for _, val := range e.many {
		if val == value {
			return true
		}
	}
	return false
}
