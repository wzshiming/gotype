package gotype

func newTypePtr(elem Type) Type {
	return &typePtr{
		Type: elem,
	}
}

type typePtr struct {
	Type
}

func (y *typePtr) Kind() Kind {
	return Ptr
}

func (t *typePtr) Elem() Type {
	return t.Type
}
