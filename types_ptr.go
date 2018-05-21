package gotype

func NewTypePtr(elem Type) Type {
	return &TypePtr{
		elem: elem,
	}
}

type TypePtr struct {
	typeBase
	elem Type
}

func (y *TypePtr) Kind() Kind {
	return Ptr
}

func (t *TypePtr) Elem() Type {
	return t.elem
}
