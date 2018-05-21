package gotype

type TypeInterface struct {
	typeBase
	methods Types // 这个类型的方法集合
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
	return t.methods.Search(name)
}
