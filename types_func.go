package gotype

type TypeFunc struct {
	typeBase
	params  []Type
	results []Type
}

func (t *TypeFunc) Kind() Kind {
	return Func
}

func (t *TypeFunc) NumOut() int {
	return len(t.results)
}

func (t *TypeFunc) Out(i int) Type {
	return t.results[i]
}

func (t *TypeFunc) NumIn() int {
	return len(t.params)
}

func (t *TypeFunc) In(i int) Type {
	return t.params[i]
}
