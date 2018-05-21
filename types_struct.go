package gotype

type TypeStruct struct {
	typeBase
	fields Types // 字段
}

func (t *TypeStruct) Kind() Kind {
	return Struct
}

func (t *TypeStruct) NumField() int {
	return t.fields.Len()
}

func (t *TypeStruct) Field(i int) Type {
	return t.fields.Index(i)
}

func (t *TypeStruct) FieldByName(name string) Type {
	return t.fields.Search(name)
}
