package gotype

func newTypeBuiltin(kind Kind) Type {
	return &typeBuiltin{
		kind: kind,
	}
}

type typeBuiltin struct {
	typeBase
	kind Kind
}

func (t *typeBuiltin) Name() string {
	return t.kind.String()
}

func (t *typeBuiltin) Kind() Kind {
	return t.kind
}
