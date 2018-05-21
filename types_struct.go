package gotype

type TypeStruct struct {
	typeBase
	fields Types // 字段
}

func (t *TypeStruct) Kind() Kind {
	return Struct
}

func (t *TypeStruct) NumField() int {
	return len(t.fields)
}

func (t *TypeStruct) Field(i int) Type {
	return t.fields[i]
}
