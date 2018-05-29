package gotype

import (
	"strings"
)

func newTypeBuiltin(kind Kind, value string) Type {
	return &typeBuiltin{
		kind:  kind,
		value: value,
	}
}

type typeBuiltin struct {
	typeBase
	kind  Kind
	value string
}

func (t *typeBuiltin) String() string {
	return strings.ToLower(t.kind.String())
}

func (t *typeBuiltin) Name() string {
	return strings.ToLower(t.kind.String())
}

func (t *typeBuiltin) Kind() Kind {
	return t.kind
}

func (t *typeBuiltin) Value() string {
	return t.value
}
