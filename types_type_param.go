package gotype

func newTypeParam(name string, constraint Type) Type {
	return &typeTypeParam{
		name:       name,
		constraint: constraint,
	}
}

type typeTypeParam struct {
	typeBase
	name       string
	constraint Type
}

func (t *typeTypeParam) Name() string {
	return t.name
}

func (t *typeTypeParam) String() string {
	if t.constraint != nil {
		return t.name + " " + t.constraint.String()
	}
	return t.name
}

func (t *typeTypeParam) Kind() Kind {
	return TypeParam
}

func (t *typeTypeParam) Constraint() Type {
	return t.constraint
}
