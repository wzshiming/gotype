package gotype

import (
	"fmt"
	"reflect"
)

type typeStructField struct {
	typeBase
	name      string
	elem      Type
	tag       reflect.StructTag
	anonymous bool
}

func (t *typeStructField) String() string {
	return fmt.Sprintf("%v %v `%v`", t.name, t.elem, t.tag)
}

func (t *typeStructField) Name() string {
	return t.name
}

func (t *typeStructField) Elem() Type {
	return t.elem
}

func (t *typeStructField) Kind() Kind {
	return Field
}

func (t *typeStructField) Tag() reflect.StructTag {
	return t.tag
}

func (t *typeStructField) IsAnonymous() bool {
	return t.anonymous
}
