package gotype

type typeStruct struct {
	typeBase
	fields  Types // 字段
	anonymo Types // 组合的类型
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

func (t *typeStruct) FieldByName(name string) Type {
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

func (t *typeStruct) MethodsByName(name string) Type {
	for _, v := range t.anonymo {
		b := v.MethodsByName(name)
		if b != nil {
			return b
		}
	}
	return nil
}

func (t *typeStruct) NumAnonymo() int {
	return t.anonymo.Len()
}

func (t *typeStruct) Anonymo(i int) Type {
	return t.anonymo.Index(i)
}

func (t *typeStruct) AnonymoByName(name string) Type {
	return t.anonymo.Search(name)
}
