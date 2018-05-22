package gotype

type TypeInterface struct {
	typeBase
	methods Types // 这个类型的方法集合
	anonymo Types // 组合的接口
}

func (t *TypeInterface) Kind() Kind {
	return Interface
}

func (t *TypeInterface) NumMethods() int {
	return t.methods.Len()
}

func (t *TypeInterface) Methods(i int) Type {
	return t.methods.Index(i)
}

func (t *TypeInterface) MethodsByName(name string) Type {
	b := t.methods.Search(name)
	if b != nil {
		return b
	}
	for _, v := range t.anonymo {
		b = v.MethodsByName(name)
		if b != nil {
			return b
		}
	}
	return nil
}

func (t *TypeInterface) NumAnonymo() int {
	return t.anonymo.Len()
}

func (t *TypeInterface) Anonymo(i int) Type {
	return t.anonymo.Index(i)
}

func (t *TypeInterface) AnonymoByName(name string) Type {
	return t.anonymo.Search(name)
}
