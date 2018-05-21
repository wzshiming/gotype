package gotype

type TypeInterface struct {
	typeBase
	methods Types // 这个类型的方法集合
}

func (t *TypeInterface) Kind() Kind {
	return Interface
}

func (t *TypeInterface) NumMethods() int {
	return len(t.methods)
}

func (t *TypeInterface) Methods(i int) Type {
	return t.methods[i]
}
