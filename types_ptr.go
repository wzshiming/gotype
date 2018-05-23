package gotype

func newTypePtr(elem Type) Type {
	return &typePtr{
		typ: elem,
	}
}

type typePtr struct {
	typeBase
	typ Type
}

func (y *typePtr) Kind() Kind {
	return Ptr
}

func (t *typePtr) Elem() Type {
	return t.typ
}
