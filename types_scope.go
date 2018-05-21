package gotype

func NewTypeScope(name string, parser *Parser) Type {
	return &TypeScope{
		name:   name,
		parser: parser,
	}
}

type TypeScope struct {
	typeBase
	name   string
	parser *Parser
}

func (t *TypeScope) ChildByName(name string) Type {
	return t.parser.Search(name)
}

func (t *TypeScope) Child(i int) Type {
	return t.parser.Child(i)
}

func (t *TypeScope) NumChild() int {
	return t.parser.NumChild()
}

func (t *TypeScope) Name() string {
	return t.name
}

func (t *TypeScope) Kind() Kind {
	return Scope
}
