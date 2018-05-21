package gotype

func NewTypeVar(name string, typ Type) *TypeVar {
	return &TypeVar{
		name: name,
		Type: typ,
	}
}

type TypeVar struct {
	Type
	name string
}

func (t *TypeVar) Name() string {
	return t.name
}

func (t *TypeVar) Kind() Kind {
	return Var
}

func (t *TypeVar) Elem() Type {
	return t.Type
}
