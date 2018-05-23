package gotype

import (
	"strings"
)

func newTypeBuiltin(kind Kind) Type {
	return &typeBuiltin{
		kind: kind,
	}
}

type typeBuiltin struct {
	typeBase
	kind Kind
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
