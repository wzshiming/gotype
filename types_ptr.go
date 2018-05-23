package gotype

import "fmt"

func newTypePtr(elem Type) Type {
	return &typePtr{
		typ: elem,
	}
}

type typePtr struct {
	typeBase
	typ Type
}

func (t *typePtr) String() string {
	return fmt.Sprintf("*%v", t.typ)
}

func (y *typePtr) Kind() Kind {
	return Ptr
}

func (t *typePtr) Elem() Type {
	return t.typ
}
