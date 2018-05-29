package gotype

import "fmt"

func newTypeArray(v Type, l int) Type {
	return &typeArray{
		elem: v,
		le:   l,
	}
}

type typeArray struct {
	typeBase
	le   int
	elem Type
}

func (t *typeArray) String() string {
	return fmt.Sprintf("[%v]%v", t.le, t.elem)
}

func (t *typeArray) Kind() Kind {
	return Array
}

func (t *typeArray) Elem() Type {
	return t.elem
}

func (t *typeArray) Len() int {
	return t.le
}
