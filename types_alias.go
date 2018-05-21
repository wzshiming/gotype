package gotype

func NewTypeAlias(name string, typ Type) Type {
	return &TypeAlias{
		name: name,
		Type: typ,
	}
}

type TypeAlias struct {
	name string
	Type
}

func (t *TypeAlias) Name() string {
	return t.name
}
