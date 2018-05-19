package gotype

type Type interface {
	Name() string
	Kind() Kind
	Key() Type
	Elem() Type // 指针 接口 别名 数组 map值 到下一级
	NumMethods() int
	Methods(int) *TypeMethod
	NumField() int
	Field(int) *TypeStructField
	Len() int     // 数组
	NumOut() int  // func
	Out(int) Type // func
	NumIn() int   // func
	In(int) Type  // func
}

type typeBase struct {
	parser *Parser
	name   string
	kind   Kind
}

func (t *typeBase) Name() string {
	return t.name
}

func (t *typeBase) Kind() Kind {
	return t.kind
}

func (typeBase) Key() Type                  { return nil }
func (typeBase) Elem() Type                 { return nil }
func (typeBase) NumField() int              { return 0 }
func (typeBase) Field(int) *TypeStructField { return nil }
func (typeBase) Len() int                   { return 0 }
func (typeBase) NumOut() int                { return 0 }
func (typeBase) Out(int) Type               { return nil }
func (typeBase) NumIn() int                 { return 0 }
func (typeBase) In(int) Type                { return nil }

func (t *typeBase) NumMethods() int {
	if t.parser == nil {
		return 0
	}
	return len(t.parser.Method[t.name])
}

func (t *typeBase) Methods(i int) *TypeMethod {
	if t.parser == nil {
		return nil
	}

	b := t.parser.Method[t.name]
	if len(b) <= i {
		return nil
	}
	return b[i]
}
