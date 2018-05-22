package gotype

func newTypeVar(name string, typ Type) *typeVar {
	v, ok := typ.(*typeVar)
	if ok {
		return &typeVar{
			name: name,
			Type: v.Type,
		}
	}
	return &typeVar{
		name: name,
		Type: typ,
	}
}

type typeVar struct {
	Type
	name string
}

func (t *typeVar) Name() string {
	return t.name
}

func (t *typeVar) Kind() Kind {
	return Var
}

func (t *typeVar) Elem() Type {
	return t.Type
}
