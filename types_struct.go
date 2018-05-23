package gotype

import (
	"bytes"
)

type typeStruct struct {
	typeBase
	fields  Types // 字段
	anonymo Types // 组合的类型
}

func (t *typeStruct) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("struct {")
	if len(t.anonymo)+len(t.fields) != 0 {
		buf.WriteByte('\n')
	}
	for _, v := range t.anonymo {
		buf.WriteString(v.String())
		buf.WriteByte('\n')
	}
	for _, v := range t.fields {
		buf.WriteString(v.String())
		buf.WriteByte('\n')
	}
	buf.WriteString("}\n")
	return buf.String()
}

func (t *typeStruct) Kind() Kind {
	return Struct
}

func (t *typeStruct) NumField() int {
	return t.fields.Len()
}

func (t *typeStruct) Field(i int) Type {
	return t.fields.Index(i)
}

func (t *typeStruct) FieldByName(name string) (Type, bool) {
	b, ok := t.fields.Search(name)
	if ok {
		return b, true
	}
	b, ok = t.anonymo.Search(name)
	if ok {
		return b, true
	}

	for _, v := range t.anonymo {
		b, ok = v.FieldByName(name)
		if ok {
			return b, true
		}
	}
	return nil, false
}

func (t *typeStruct) MethodsByName(name string) (Type, bool) {
	for _, v := range t.anonymo {
		b, ok := v.MethodsByName(name)
		if ok {
			return b, true
		}
	}
	return nil, false
}

func (t *typeStruct) NumAnonymo() int {
	return t.anonymo.Len()
}

func (t *typeStruct) Anonymo(i int) Type {
	return t.anonymo.Index(i)
}

func (t *typeStruct) AnonymoByName(name string) (Type, bool) {
	return t.anonymo.Search(name)
}
