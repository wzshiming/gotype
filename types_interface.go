package gotype

type typeInterface struct {
	typeBase
	methods Types // 这个类型的方法集合
	anonymo Types // 组合的接口
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

func (t *typeInterface) MethodsByName(name string) Type {
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

func (t *typeInterface) NumAnonymo() int {
	return t.anonymo.Len()
}

func (t *typeInterface) Anonymo(i int) Type {
	return t.anonymo.Index(i)
}

func (t *typeInterface) AnonymoByName(name string) Type {
	return t.anonymo.Search(name)
}
