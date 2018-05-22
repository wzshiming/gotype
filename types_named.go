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
	child := t.ToChild()
	if child == nil {
		return Invalid
	}
	return child.Kind()
}

func (t *TypeNamed) Key() Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.Key()
}

func (t *TypeNamed) Elem() Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.Elem()
}

func (t *TypeNamed) NumField() int {
	child := t.ToChild()
	if child == nil {
		return 0
	}
	return child.NumField()
}

func (t *TypeNamed) Field(i int) Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.Field(i)
}

func (t *TypeNamed) FieldByName(name string) Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.FieldByName(name)
}

func (t *TypeNamed) Len() int {
	child := t.ToChild()
	if child == nil {
		return 0
	}
	return child.Len()
}

func (t *TypeNamed) NumOut() int {
	child := t.ToChild()
	if child == nil {
		return 0
	}
	return child.NumOut()
}

func (t *TypeNamed) Out(i int) Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.Out(i)
}

func (t *TypeNamed) NumIn() int {
	child := t.ToChild()
	if child == nil {
		return 0
	}
	return child.NumIn()
}

func (t *TypeNamed) In(i int) Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.In(i)
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
