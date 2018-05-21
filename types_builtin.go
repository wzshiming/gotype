package gotype

func NewTypeBuiltin(kind Kind) Type {
	return &TypeBuiltin{
		kind: kind,
	}
}

type TypeBuiltin struct {
	typeBase
	kind Kind
}

func (t *TypeBuiltin) Name() string {
	return t.kind.String()
}

func (t *TypeBuiltin) Kind() Kind {
	return t.kind
}
