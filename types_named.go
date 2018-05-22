package gotype

func NewTypeNamed(name string, typ Type, parser *Parser) Type {
	return &TypeNamed{
		name:   name,
		Type:   typ,
		parser: parser,
	}
}

type TypeNamed struct {
	name   string
	parser *Parser
	Type
}

func (t *TypeNamed) ToChild() Type {
	if t.Type == nil {
		t.Type = t.parser.Search(t.Name())
	}
	return t.Type
}

func (t *TypeNamed) Name() string {
	return t.name
}

func (t *TypeNamed) Kind() Kind {
	return t.ToChild().Kind()
}

func (t *TypeNamed) Key() Type {
	return t.ToChild().Key()
}

func (t *TypeNamed) Elem() Type {
	return t.ToChild()
}

func (t *TypeNamed) NumField() int {
	return t.ToChild().NumField()
}

func (t *TypeNamed) Field(i int) Type {
	return t.ToChild().Field(i)
}

func (t *TypeNamed) FieldByName(name string) Type {
	return t.ToChild().FieldByName(name)
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
	if t.parser == nil {
		return 0
	}
	b := t.parser.Method[t.Name()]
	return b.Len()
}

func (t *TypeNamed) Methods(i int) Type {
	if t.parser == nil {
		return nil
	}
	b := t.parser.Method[t.Name()]
	if b.Len() <= i {
		return nil
	}
	return b.Index(i)
}

func (t *TypeNamed) MethodsByName(name string) Type {
	if t.parser == nil {
		return nil
	}
	b := t.parser.Method[t.Name()]
	return b.Search(name)
}
