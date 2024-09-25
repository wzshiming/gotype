package gotype

type typeParam struct {
	typeBase
	name string
	elem Type
}

func (t *typeParam) String() string {
	if t.name != "" {
		return t.name
	}
	return t.elem.String()
}

func (t *typeParam) Name() string {
	return t.name
}

func (t *typeParam) Elem() Type {
	return t.elem
}

func (t *typeParam) Kind() Kind {
	return Param
}
