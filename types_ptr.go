package gotype

type TypePtr struct {
	TypeBuiltin
	elem Type
}

func (t *TypePtr) Elem() Type {
	return t.elem
}
