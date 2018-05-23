package gotype

import (
	"fmt"
)

func newTypeSlice(v Type) Type {
	return &typeSlice{
		val: v,
	}
}

type typeSlice struct {
	typeBase
	val Type
}

func (t *typeSlice) String() string {
	return fmt.Sprintf("[]%v", t.val.String())
}

func (t *typeSlice) Kind() Kind {
	return Slice
}

func (t *typeSlice) Elem() Type {
	return t.val
}
