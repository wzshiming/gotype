package gotype

// newTypeGeneric creates a wrapper around a type that adds type parameters
func newTypeGeneric(base Type, typeParams []Type) Type {
	if len(typeParams) == 0 {
		return base
	}
	return &typeGeneric{
		Type:       base,
		typeParams: typeParams,
	}
}

type typeGeneric struct {
	Type
	typeParams []Type
}

func (t *typeGeneric) NumTypeParam() int {
	return len(t.typeParams)
}

func (t *typeGeneric) TypeParam(i int) Type {
	if i < 0 || i >= len(t.typeParams) {
		panic("TypeParam index out of range")
	}
	return t.typeParams[i]
}

func (t *typeGeneric) String() string {
	if len(t.typeParams) == 0 {
		return t.Type.String()
	}
	
	s := t.Type.String()
	s += "["
	for i, param := range t.typeParams {
		if i > 0 {
			s += ", "
		}
		s += param.String()
	}
	s += "]"
	return s
}
