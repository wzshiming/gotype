package gotype

import (
	"fmt"
	"reflect"
)

func newSelector(x Type, sel string) *typeSelector {
	return &typeSelector{
		x:   x,
		sel: sel,
	}
}

type typeSelector struct {
	x   Type
	sel string
	typ Type
}

func (t *typeSelector) ToChild() (Type, bool) {
	if t.typ != nil {
		return t.typ, true
	}
	s := t.x
	for {
		k := s.Kind()
		if k != Var && k != Ptr {
			break
		}
		s = s.Elem()
	}
	name := t.sel

	b, ok := s.ChildByName(name)
	if ok {
		t.typ = b
		return b, true
	}
	b, ok = s.MethodsByName(name)
	if ok {
		t.typ = b
		return b, true
	}
	if s.Kind() == Struct {
		b, ok = s.FieldByName(name)
		if ok {
			t.typ = b
			return b, true
		}
	}
	return nil, false
}

func (t *typeSelector) String() string {
	return fmt.Sprintf("%v.%v", t.x, t.sel)
}

func (t *typeSelector) Name() string {
	return t.sel
}

func (t *typeSelector) Kind() Kind {
	child, ok := t.ToChild()
	if !ok {
		return Invalid
	}
	return child.Kind()
}

func (t *typeSelector) Key() Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Key()
}

func (t *typeSelector) Elem() Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Elem()
}

func (t *typeSelector) NumField() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumField()
}

func (t *typeSelector) Field(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Field(i)
}

func (t *typeSelector) FieldByName(name string) (Type, bool) {
	child, ok := t.ToChild()
	if !ok {
		return nil, false
	}
	return child.FieldByName(name)
}

func (t *typeSelector) Tag() reflect.StructTag {
	child, ok := t.ToChild()
	if !ok {
		return ""
	}
	return child.Tag()
}

func (t *typeSelector) Len() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.Len()
}

func (t *typeSelector) ChanDir() ChanDir {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.ChanDir()
}

func (t *typeSelector) NumOut() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumOut()
}

func (t *typeSelector) Out(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Out(i)
}

func (t *typeSelector) NumIn() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumIn()
}

func (t *typeSelector) In(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.In(i)
}

func (t *typeSelector) IsVariadic() bool {
	child, ok := t.ToChild()
	if !ok {
		return false
	}
	return child.IsVariadic()
}

func (t *typeSelector) NumMethods() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumMethods()
}

func (t *typeSelector) Methods(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Methods(i)
}

func (t *typeSelector) MethodsByName(name string) (Type, bool) {
	child, ok := t.ToChild()
	if !ok {
		return nil, false
	}
	return child.MethodsByName(name)
}

func (t *typeSelector) Child(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Child(i)
}

func (t *typeSelector) ChildByName(name string) (Type, bool) {
	child, ok := t.ToChild()
	if !ok {
		return nil, false
	}
	return child.ChildByName(name)
}

func (t *typeSelector) NumChild() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumChild()
}

func (t *typeSelector) Anonymo(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Anonymo(i)
}

func (t *typeSelector) AnonymoByName(name string) (Type, bool) {
	child, ok := t.ToChild()
	if !ok {
		return nil, false
	}
	return child.AnonymoByName(name)
}

func (t *typeSelector) NumAnonymo() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumAnonymo()
}
