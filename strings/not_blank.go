package strings

import (
	"github.com/begopher/valid"
	"strings"
)

func NotBlank() valid.Predicate[string] {
	return notBlank{}
}

type notBlank struct{}

func (notBlank) Evaluate(value string) bool {
	if "" != strings.TrimSpace(value) {
		return true
	}
	return false
}

