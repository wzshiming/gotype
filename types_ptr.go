package gotype

func newTypePtr(elem Type) Type {
	return &typePtr{
		elem: elem,
	}
}

type typePtr struct {
	typeBase
	elem Type
}

func (y *typePtr) Kind() Kind {
	return Ptr
}

func (t *typePtr) Elem() Type {
	return t.elem
}
