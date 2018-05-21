package gotype

type TypeNamed struct {
	TypeBuiltin
	typ         Type
	resetMethod bool // 是否依赖指向的 方法
}

func (t *TypeNamed) ToChild() Type {
	if t.typ == nil {
		t.typ = t.parser.ChildByName(t.name)
	}
	return t.typ
}

func (t *TypeNamed) Kind() Kind {
	return t.ToChild().Kind()
}

func (t *TypeNamed) Key() Type {
	return t.ToChild().Key()
}

func (t *TypeNamed) Elem() Type {
	return t.ToChild().Elem()
}

func (t *TypeNamed) NumField() int {
	return t.ToChild().NumField()
}

func (t *TypeNamed) Field(i int) *TypeStructField {
	return t.ToChild().Field(i)
}

func (t *TypeNamed) Len() int {
	return t.ToChild().Len()
}

func (t *TypeNamed) NumOut() int {
	return t.ToChild().NumOut()
}

func (t *TypeNamed) Out(i int) Type {
	return t.ToChild().Out(i)
}

func (t *TypeNamed) NumIn() int {
	return t.ToChild().NumIn()
}

func (t *TypeNamed) In(i int) Type {
	return t.ToChild().In(i)
}

func (t *TypeNamed) NumMethods() int {
	if t.resetMethod {
		return t.TypeBuiltin.NumMethods()
	}
	return t.ToChild().NumMethods()
}

func (t *TypeNamed) Methods(i int) *TypeMethod {
	if t.resetMethod {
		return t.TypeBuiltin.Methods(i)
	}
	return t.ToChild().Methods(i)
}
