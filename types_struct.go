package gotype

type TypeStruct struct {
	typeBase
	fields  Types // 字段
	anonymo Types // 组合的类型
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
	b := t.fields.Search(name)
	if b != nil {
		return b
	}
	b = t.anonymo.Search(name)
	if b != nil {
		return b
	}

	for _, v := range t.anonymo {
		b = v.FieldByName(name)
		if b != nil {
			return b
		}
	}
	return nil
}

func (t *TypeStruct) MethodsByName(name string) Type {
	for _, v := range t.anonymo {
		b := v.MethodsByName(name)
		if b != nil {
			return b
		}
	}
	return nil
}

func (t *TypeStruct) NumAnonymo() int {
	return t.anonymo.Len()
}

func (t *TypeStruct) Anonymo(i int) Type {
	return t.anonymo.Index(i)
}

func (t *TypeStruct) AnonymoByName(name string) Type {
	return t.anonymo.Search(name)
}
