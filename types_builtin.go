package gotype

var _ Type = (*TypeBuiltin)(nil)

type TypeBuiltin struct {
	typeBase
}

func (t *TypeBuiltin) Name() string {
	return t.kind.String()
}
