package gotype

type TypePtr struct {
	typeBase
	elem Type
}

func (t *TypePtr) Elem() Type {
	return t.elem
}
