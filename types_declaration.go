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
