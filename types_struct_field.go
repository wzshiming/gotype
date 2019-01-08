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
	tag := ""
	if t.tag != "" {
		tag = " `" + string(t.tag) + "`"
	}
	if t.anonymous {
		return fmt.Sprintf("%v%s", t.elem, tag)
	}
	return fmt.Sprintf("%v %v%s", t.name, t.elem, tag)
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
