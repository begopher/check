package strings

import (
	"github.com/begopher/check"
	"unicode/utf8"
)

func More(runesNum int) check.Predicate[string] {
	return more{runesNum}
}

type more struct {
	number int
}

func (l more) Evaluate(value string) bool {
	n := utf8.RuneCountInString(value)
	if n > l.number {
		return true
	}
	return false
}
