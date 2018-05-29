package gotype

import (
	"fmt"
)

func newTypeSlice(v Type) Type {
	return &typeSlice{
		elem: v,
	}
}

type typeSlice struct {
	typeBase
	elem Type
}

func (t *typeSlice) String() string {
	return fmt.Sprintf("[]%v", t.elem.String())
}

func (t *typeSlice) Kind() Kind {
	return Slice
}

func (t *typeSlice) Elem() Type {
	return t.elem
}
