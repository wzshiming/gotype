package gotype

func newTypeScope(name string, info *info) Type {
	return &typeScope{
		name: name,
		info: info,
	}
}

type typeScope struct {
	typeBase
	name string
	info *info
}

func (t *typeScope) String() string {
	return t.name
}

func (t *typeScope) ChildByName(name string) (Type, bool) {
	return t.info.Named.Search(name)
}

func (t *typeScope) Child(i int) Type {
	return t.info.Named.Index(i)
}

func (t *typeScope) NumChild() int {
	return t.info.Named.Len()
}

func (t *typeScope) Name() string {
	return t.name
}

func (t *typeScope) Kind() Kind {
	return Scope
}
