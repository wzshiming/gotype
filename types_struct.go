package gotype

import "reflect"

type TypeStructField struct {
	Name string
	Type Type              // 字段类型
	Tag  reflect.StructTag // 字段标签
}

type TypeStruct struct {
	TypeBuiltin
	name   string             // 名字
	fields []*TypeStructField // 字段
}

func (t *TypeStruct) Name() string {
	return t.name
}

func (t *TypeStruct) NumField() int {
	return len(t.fields)
}

func (t *TypeStruct) Field(i int) *TypeStructField {
	return t.fields[i]
}
