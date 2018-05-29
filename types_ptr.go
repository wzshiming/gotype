package gotype

import "fmt"

func newTypePtr(elem Type) Type {
	return &typePtr{
		elem: elem,
	}
}

type typePtr struct {
	typeBase
	elem Type
}

func (t *typePtr) String() string {
	return fmt.Sprintf("*%v", t.elem)
}

func (y *typePtr) Kind() Kind {
	return Ptr
}

func (t *typePtr) Elem() Type {
	return t.elem
}
