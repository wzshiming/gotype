package gotype

type TypeFunc struct {
	typeBase
	params  Types
	results Types
}

func (t *TypeFunc) Kind() Kind {
	return Func
}

func (t *TypeFunc) NumOut() int {
	return t.results.Len()
}

func (t *TypeFunc) Out(i int) Type {
	return t.results.Index(i)
}

func (t *TypeFunc) NumIn() int {
	return t.params.Len()
}

func (t *TypeFunc) In(i int) Type {
	return t.params.Index(i)
}
