package strings

import (
	"github.com/begopher/check"
	"strings"
)

func Blank() check.Predicate[string] {
	return blank{}
}

type blank struct{}

func (blank) Evaluate(value string) bool {
	if value == strings.TrimSpace(value) {
		return true
	}
	return false
}
