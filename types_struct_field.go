package gotype

import (
	"fmt"
	"reflect"
)

type typeStructField struct {
	typeBase
	name string
	typ  Type              // 字段类型
	tag  reflect.StructTag // 字段标签
}

func (t *typeStructField) String() string {
	return fmt.Sprintf("%v %v %v", t.name, t.typ, t.tag)
}

func (t *typeStructField) Name() string {
	return t.name
}

func (t *typeStructField) Elem() Type {
	return t.typ
}

func (t *typeStructField) Kind() Kind {
	return Field
}

func (t *typeStructField) Tag() reflect.StructTag {
	return t.tag
}
