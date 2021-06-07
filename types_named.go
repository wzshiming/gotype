package gotype

import (
	"go/ast"
	"reflect"
)

func newTypeNamed(name string, typ Type, info *infoFile) Type {
	return &typeNamed{
		name: name,
		typ:  typ,
		info: info,
	}
}

type typeNamed struct {
	name string
	info *infoFile
	typ  Type
}

func (t *typeNamed) ToChild() (Type, bool) {
	if t.typ == nil {
		var ok bool
		t.typ, ok = t.info.GetPkgOrType(t.Name())
		return t.typ, ok
	}
	return t.typ, true
}

func (t *typeNamed) PkgPath() string {
	child, ok := t.ToChild()
	if !ok {
		return ""
	}
	return child.PkgPath()
}

func (t *typeNamed) IsGoroot() bool {
	child, ok := t.ToChild()
	if !ok {
		return false
	}
	return child.IsGoroot()
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

func (t *typeNamed) Declaration() Type {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Declaration()
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

func (t *typeNamed) IsAnonymous() bool {
	child, ok := t.ToChild()
	if !ok {
		return false
	}
	return child.IsAnonymous()
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

func (t *typeNamed) Value() string {
	child, ok := t.ToChild()
	if !ok {
		return ""
	}
	return child.Value()
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

func (t *typeNamed) NumMethod() int {
	if t.info == nil {
		return 0
	}
	b := t.info.Methods[t.Name()]
	return b.Len()
}

func (t *typeNamed) Method(i int) Type {
	if t.info == nil {
		return nil
	}
	b := t.info.Methods[t.Name()]
	if b.Len() <= i {
		return nil
	}
	return b.Index(i)
}

func (t *typeNamed) MethodByName(name string) (Type, bool) {
	if t.info == nil {
		return nil, false
	}
	b := t.info.Methods[t.Name()]
	m, ok := b.Search(name)
	if ok {
		return m, true
	}
	child, ok := t.ToChild()
	if ok {
		return child.MethodByName(name)
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

func (t *typeNamed) Origin() ast.Node {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Origin()
}

func (t *typeNamed) Doc() *ast.CommentGroup {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Doc()
}

func (t *typeNamed) Comment() *ast.CommentGroup {
	child, ok := t.ToChild()
	if !ok {
		return nil
	}
	return child.Comment()
}
