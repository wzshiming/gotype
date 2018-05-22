package gotype

func newTypeAlias(name string, typ Type) Type {
	return &typeAlias{
		name: name,
		Type: typ,
	}
}

type typeAlias struct {
	name string
	Type
}

func (t *typeAlias) Name() string {
	return t.name
}
