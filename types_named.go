package gotype

import (
	"reflect"
)

func newTypeNamed(name string, typ Type, parser *parser) Type {
	return &typeNamed{
		name:   name,
		typ:    typ,
		parser: parser,
	}
}

type typeNamed struct {
	name   string
	parser *parser
	typ    Type
}

func (t *typeNamed) ToChild() (Type, bool) {
	if t.typ == nil {
		var ok bool
		t.typ, ok = t.parser.nameds.Search(t.Name())
		return t.typ, ok
	}
	return t.typ, true
}

func (t *typeNamed) Name() string {
	return t.name
}

func (t *typeNamed) String() string {
	return t.name
}

func (t *typeNamed) Kind() Kind {
	child, ok := t.ToChild()
	if !ok {
		return Invalid
	}
	return child.Kind()
}

func (t *typeNamed) Key() Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Key()
}

func (t *typeNamed) Elem() Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Elem()
}

func (t *typeNamed) NumField() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumField()
}

func (t *typeNamed) Field(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Field(i)
}

func (t *typeNamed) FieldByName(name string) (Type, bool) {
	child, ok := t.ToChild()
	if !ok {
		return nil, false
	}
	return child.FieldByName(name)
}

func (t *typeNamed) Tag() reflect.StructTag {
	child, ok := t.ToChild()
	if !ok {
		return ""
	}
	return child.Tag()
}

func (t *typeNamed) Len() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.Len()
}

func (t *typeNamed) ChanDir() ChanDir {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.ChanDir()
}

func (t *typeNamed) NumOut() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumOut()
}

func (t *typeNamed) Out(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Out(i)
}

func (t *typeNamed) NumIn() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumIn()
}

func (t *typeNamed) In(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.In(i)
}

func (t *typeNamed) IsVariadic() bool {
	child, ok := t.ToChild()
	if !ok {
		return false
	}
	return child.IsVariadic()
}

func (t *typeNamed) NumMethods() int {
	if t.parser == nil {
		return 0
	}
	b := t.parser.method[t.Name()]
	return b.Len()
}

func (t *typeNamed) Methods(i int) Type {
	if t.parser == nil {
		return nil
	}
	b := t.parser.method[t.Name()]
	if b.Len() <= i {
		return nil
	}
	return b.Index(i)
}

func (t *typeNamed) MethodsByName(name string) (Type, bool) {
	if t.parser == nil {
		return nil, false
	}
	b := t.parser.method[t.Name()]
	m, ok := b.Search(name)
	if ok {
		return m, true
	}
	child, ok := t.ToChild()
	if ok {
		return child.MethodsByName(name)
	}
	return nil, false
}

func (t *typeNamed) Child(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Child(i)
}

func (t *typeNamed) ChildByName(name string) (Type, bool) {
	child, ok := t.ToChild()
	if !ok {
		return nil, false
	}
	return child.ChildByName(name)
}

func (t *typeNamed) NumChild() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumChild()
}

func (t *typeNamed) Anonymo(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Anonymo(i)
}

func (t *typeNamed) AnonymoByName(name string) (Type, bool) {
	child, ok := t.ToChild()
	if !ok {
		return nil, false
	}
	return child.AnonymoByName(name)
}

func (t *typeNamed) NumAnonymo() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumAnonymo()
}
