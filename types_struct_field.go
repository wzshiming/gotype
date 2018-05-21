package gotype

import (
	"reflect"
)

type TypeStructField struct {
	typeBase
	name string
	typ  Type              // 字段类型
	tag  reflect.StructTag // 字段标签
}

func (t *TypeStructField) Name() string {
	return t.name
}

func (t *TypeStructField) Elem() Type {
	return t.typ
}

func (t *TypeStructField) Kind() Kind {
	return Field
}

func (t *TypeStructField) Tag() reflect.StructTag {
	return t.tag
}
