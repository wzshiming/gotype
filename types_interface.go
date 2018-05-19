package gotype

type TypeMethod struct {
	Name string
	Func Type
}

type TypeInterface struct {
	typeBase
	name    string        // 名字
	methods []*TypeMethod // 这个类型的方法集合
}

func (t *TypeInterface) Name() string {
	return t.name
}

func (t *TypeInterface) NumMethods() int {
	return len(t.methods)
}

func (t *TypeInterface) Methods(i int) *TypeMethod {
	return t.methods[i]
}
