package gotype

func newTypeNamed(name string, typ Type, parser *astParser) Type {
	return &typeNamed{
		name:   name,
		Type:   typ,
		parser: parser,
	}
}

type typeNamed struct {
	name   string
	parser *astParser
	Type
}

func (t *typeNamed) ToChild() Type {
	if t.Type == nil {
		t.Type = t.parser.Search(t.Name())
	}
	return t.Type
}

func (t *typeNamed) Name() string {
	return t.name
}

func (t *typeNamed) Kind() Kind {
	child := t.ToChild()
	if child == nil {
		return Invalid
	}
	return child.Kind()
}

func (t *typeNamed) Key() Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.Key()
}

func (t *typeNamed) Elem() Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.Elem()
}

func (t *typeNamed) NumField() int {
	child := t.ToChild()
	if child == nil {
		return 0
	}
	return child.NumField()
}

func (t *typeNamed) Field(i int) Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.Field(i)
}

func (t *typeNamed) FieldByName(name string) Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.FieldByName(name)
}

func (t *typeNamed) Len() int {
	child := t.ToChild()
	if child == nil {
		return 0
	}
	return child.Len()
}

func (t *typeNamed) NumOut() int {
	child := t.ToChild()
	if child == nil {
		return 0
	}
	return child.NumOut()
}

func (t *typeNamed) Out(i int) Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.Out(i)
}

func (t *typeNamed) NumIn() int {
	child := t.ToChild()
	if child == nil {
		return 0
	}
	return child.NumIn()
}

func (t *typeNamed) In(i int) Type {
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.In(i)
}

func (t *typeNamed) NumMethods() int {
	if t.parser == nil {
		return 0
	}
	b := t.parser.Method[t.Name()]
	return b.Len()
}

func (t *typeNamed) Methods(i int) Type {
	if t.parser == nil {
		return nil
	}
	b := t.parser.Method[t.Name()]
	if b.Len() <= i {
		return nil
	}
	return b.Index(i)
}

func (t *typeNamed) MethodsByName(name string) Type {
	if t.parser == nil {
		return nil
	}
	b := t.parser.Method[t.Name()]
	m := b.Search(name)
	if m != nil {
		return m
	}
	child := t.ToChild()
	if child == nil {
		return nil
	}
	return child.MethodsByName(name)
}
