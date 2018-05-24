package gotype

import "bytes"

type typeInterface struct {
	typeBase
	methods types // 这个类型的方法集合
	anonymo types // 组合的接口
}

func (t *typeInterface) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("interface {")
	if len(t.anonymo)+len(t.methods) != 0 {
		buf.WriteByte('\n')
	}
	for _, v := range t.anonymo {
		buf.WriteString(v.String())
		buf.WriteByte('\n')
	}
	for _, v := range t.methods {
		buf.WriteString(v.String())
		buf.WriteByte('\n')
	}
	buf.WriteString("}\n")
	return buf.String()
}

func (t *typeInterface) Kind() Kind {
	return Interface
}

func (t *typeInterface) NumMethods() int {
	return t.methods.Len()
}

func (t *typeInterface) Methods(i int) Type {
	return t.methods.Index(i)
}

func (t *typeInterface) MethodsByName(name string) (Type, bool) {
	b, ok := t.methods.Search(name)
	if ok {
		return b, true
	}
	for _, v := range t.anonymo {
		b, ok := v.MethodsByName(name)
		if ok {
			return b, true
		}
	}
	return nil, false
}

func (t *typeInterface) NumAnonymo() int {
	return t.anonymo.Len()
}

func (t *typeInterface) Anonymo(i int) Type {
	return t.anonymo.Index(i)
}

func (t *typeInterface) AnonymoByName(name string) (Type, bool) {
	return t.anonymo.Search(name)
}
