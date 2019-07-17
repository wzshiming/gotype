package gotype

import (
	"strings"
)

func newTypeBuiltin(kind Kind, value string) Type {
	return &typeBuiltin{
		kind:  kind,
		name:  strings.ToLower(kind.String()),
		value: value,
	}
}

type typeBuiltin struct {
	typeBase
	kind  Kind
	name  string
	value string
}

func (t *typeBuiltin) String() string {
	return t.name
}

func (t *typeBuiltin) Name() string {
	return t.name
}

func (t *typeBuiltin) Kind() Kind {
	return t.kind
}

func (t *typeBuiltin) Value() string {
	return t.value
}
