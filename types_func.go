package gotype

type typeFunc struct {
	typeBase
	variadic bool
	params   Types
	results  Types
}

func (t *typeFunc) Kind() Kind {
	return Func
}

func (t *typeFunc) NumOut() int {
	return t.results.Len()
}

func (t *typeFunc) Out(i int) Type {
	return t.results.Index(i)
}

func (t *typeFunc) NumIn() int {
	return t.params.Len()
}

func (t *typeFunc) In(i int) Type {
	return t.params.Index(i)
}

func (t *typeFunc) IsVariadic() bool {
	return t.variadic
}
