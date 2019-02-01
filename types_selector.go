package gotype

import (
	"fmt"
	"go/ast"
	"reflect"
)

func newSelector(x Type, sel string) Type {
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
		if k == Declaration {
			s = s.Declaration()
			continue
		}
		if k == Ptr {
			s = s.Elem()
			continue
		}
		break
	}
	name := t.sel

	b, ok := s.ChildByName(name)
	if ok {
		t.typ = b
		return b, true
	}
	b, ok = s.MethodByName(name)
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
	if t.x.Kind() == Scope {
		return t.sel
	}
	return fmt.Sprintf("%v.%v", t.x, t.sel)
}

func (t *typeSelector) PkgPath() string {
	child, ok := t.ToChild()
	if !ok {
		return ""
	}
	return child.PkgPath()
}

func (t *typeSelector) IsGoroot() bool {
	child, ok := t.ToChild()
	if !ok {
		return false
	}
	return child.IsGoroot()
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

func (t *typeSelector) Declaration() Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Declaration()
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

func (t *typeSelector) IsAnonymous() bool {
	child, ok := t.ToChild()
	if !ok {
		return false
	}
	return child.IsAnonymous()
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

func (t *typeSelector) Value() string {
	child, ok := t.ToChild()
	if !ok {
		return ""
	}
	return child.Value()
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

func (t *typeSelector) NumMethod() int {
	child, ok := t.ToChild()
	if !ok {
		return 0
	}
	return child.NumMethod()
}

func (t *typeSelector) Method(i int) Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Method(i)
}

func (t *typeSelector) MethodByName(name string) (Type, bool) {
	child, ok := t.ToChild()
	if !ok {
		return nil, false
	}
	return child.MethodByName(name)
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

func (t *typeSelector) Origin() ast.Node {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Origin()
}

func (t *typeSelector) Doc() *ast.CommentGroup {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Doc()
}

func (t *typeSelector) Comment() *ast.CommentGroup {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Comment()
}
