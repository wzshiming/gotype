package gotype

func newDeclaration(name string, typ Type) Type {
	return &typeDeclaration{
		name:        name,
		declaration: typ,
	}
}

type typeDeclaration struct {
	typeBase
	declaration Type
	name        string
}

func (t *typeDeclaration) String() string {
	if t.declaration.NumTypeParam() > 0 {
		// Build string with type parameters
		s := t.name + "["
		for i := 0; i < t.declaration.NumTypeParam(); i++ {
			if i > 0 {
				s += ", "
			}
			s += t.declaration.TypeParam(i).String()
		}
		s += "]"
		return s
	}
	return t.name
}

func (t *typeDeclaration) Name() string {
	return t.name
}

func (t *typeDeclaration) Kind() Kind {
	return Declaration
}

func (t *typeDeclaration) Declaration() Type {
	return t.declaration
}

func (t *typeDeclaration) Value() string {
	return t.declaration.Value()
}
