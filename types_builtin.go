package gotype

type TypeBuiltin struct {
	parser *Parser
	name    string
	kind    Kind
}

func (t *TypeBuiltin) Name() string {
	return t.name
}

func (t *TypeBuiltin) Kind() Kind {
	return t.kind
}

func (TypeBuiltin) Key() Type                  { return nil }
func (TypeBuiltin) Elem() Type                 { return nil }
func (TypeBuiltin) NumField() int              { return 0 }
func (TypeBuiltin) Field(int) *TypeStructField { return nil }
func (TypeBuiltin) Len() int                   { return 0 }
func (TypeBuiltin) NumOut() int                { return 0 }
func (TypeBuiltin) Out(int) Type               { return nil }
func (TypeBuiltin) NumIn() int                 { return 0 }
func (TypeBuiltin) In(int) Type                { return nil }

func (t *TypeBuiltin) NumMethods() int {
	if t.parser == nil {
		return 0
	}
	return len(t.parser.Method[t.name])
}

func (t *TypeBuiltin) Methods(i int) *TypeMethod {
	if t.parser == nil {
		return nil
	}
	b := t.parser.Method[t.name]
	if len(b) <= i {
		return nil
	}
	return b[i]
}
