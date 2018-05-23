package gotype

func newTypeScope(name string, parser *astParser) Type {
	return &typeScope{
		name:   name,
		parser: parser,
	}
}

type typeScope struct {
	typeBase
	name   string
	parser *astParser
}

func (t *typeScope) String() string {
	return t.name
}

func (t *typeScope) ChildByName(name string) (Type, bool) {
	return t.parser.nameds.Search(name)
}

func (t *typeScope) Child(i int) Type {
	return t.parser.nameds.Index(i)
}

func (t *typeScope) NumChild() int {
	return t.parser.nameds.Len()
}

func (t *typeScope) Name() string {
	return t.name
}

func (t *typeScope) Kind() Kind {
	return Scope
}
