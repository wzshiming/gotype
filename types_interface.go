package gotype

import "bytes"

type typeInterface struct {
	typeBase
	all     types
	anonymo types
	method  types
}

func (t *typeInterface) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("interface{")
	for i, v := range t.all {
		if i != 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(v.String())
	}
	buf.WriteByte('}')
	return buf.String()
}

func (t *typeInterface) Kind() Kind {
	return Interface
}

func (t *typeInterface) NumMethod() int {
	return t.method.Len()
}

func (t *typeInterface) Method(i int) Type {
	return t.method.Index(i)
}

func (t *typeInterface) MethodByName(name string) (Type, bool) {
	b, ok := t.method.Search(name)
	if ok {
		return b, true
	}
	for _, v := range t.anonymo {
		b, ok := v.MethodByName(name)
		if ok {
			return b, true
		}
	}
	return nil, false
}

func (t *typeInterface) NumField() int {
	return t.all.Len()
}

func (t *typeInterface) Field(i int) Type {
	return t.all.Index(i)
}

func (t *typeInterface) FieldByName(name string) (Type, bool) {
	b, ok := t.all.Search(name)
	if ok {
		return b, true
	}
	return nil, false
}
