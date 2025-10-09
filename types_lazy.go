package gotype

import (
	"go/ast"
	"sync"
)

// typeLazy wraps an ast.Decl and defers parsing until the type is accessed
type typeLazy struct {
	typeBase
	once     sync.Once
	parser   *parser
	info     *infoFile
	decl     ast.Decl
	parsed   Type
	name     string
	comments *ast.CommentGroup
}

func newTypeLazy(name string, decl ast.Decl, parser *parser, info *infoFile, comments *ast.CommentGroup) *typeLazy {
	return &typeLazy{
		name:     name,
		decl:     decl,
		parser:   parser,
		info:     info,
		comments: comments,
	}
}

func (t *typeLazy) parse() {
	t.once.Do(func() {
		t.parser.parseDecl(t.info, t.decl)
		// After parsing, lookup the actual parsed type
		if typ, ok := t.info.Named.Search(t.name); ok {
			t.parsed = typ
		}
	})
}

func (t *typeLazy) String() string {
	t.parse()
	if t.parsed != nil {
		return t.parsed.String()
	}
	return t.name
}

func (t *typeLazy) PkgPath() string {
	t.parse()
	if t.parsed != nil {
		return t.parsed.PkgPath()
	}
	return ""
}

func (t *typeLazy) IsGoroot() bool {
	t.parse()
	if t.parsed != nil {
		return t.parsed.IsGoroot()
	}
	return false
}

func (t *typeLazy) Name() string {
	return t.name
}

func (t *typeLazy) Kind() Kind {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Kind()
	}
	return Invalid
}

func (t *typeLazy) Key() Type {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Key()
	}
	return nil
}

func (t *typeLazy) Elem() Type {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Elem()
	}
	return nil
}

func (t *typeLazy) Declaration() Type {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Declaration()
	}
	return nil
}

func (t *typeLazy) Len() int {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Len()
	}
	return 0
}

func (t *typeLazy) Value() string {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Value()
	}
	return ""
}

func (t *typeLazy) ChanDir() ChanDir {
	t.parse()
	if t.parsed != nil {
		return t.parsed.ChanDir()
	}
	return 0
}

func (t *typeLazy) Out(i int) Type {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Out(i)
	}
	return nil
}

func (t *typeLazy) NumOut() int {
	t.parse()
	if t.parsed != nil {
		return t.parsed.NumOut()
	}
	return 0
}

func (t *typeLazy) In(i int) Type {
	t.parse()
	if t.parsed != nil {
		return t.parsed.In(i)
	}
	return nil
}

func (t *typeLazy) NumIn() int {
	t.parse()
	if t.parsed != nil {
		return t.parsed.NumIn()
	}
	return 0
}

func (t *typeLazy) IsVariadic() bool {
	t.parse()
	if t.parsed != nil {
		return t.parsed.IsVariadic()
	}
	return false
}

func (t *typeLazy) Field(i int) Type {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Field(i)
	}
	return nil
}

func (t *typeLazy) FieldByName(name string) (Type, bool) {
	t.parse()
	if t.parsed != nil {
		return t.parsed.FieldByName(name)
	}
	return nil, false
}

func (t *typeLazy) NumField() int {
	t.parse()
	if t.parsed != nil {
		return t.parsed.NumField()
	}
	return 0
}

func (t *typeLazy) IsAnonymous() bool {
	t.parse()
	if t.parsed != nil {
		return t.parsed.IsAnonymous()
	}
	return false
}

func (t *typeLazy) Method(i int) Type {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Method(i)
	}
	return nil
}

func (t *typeLazy) MethodByName(name string) (Type, bool) {
	t.parse()
	if t.parsed != nil {
		return t.parsed.MethodByName(name)
	}
	return nil, false
}

func (t *typeLazy) NumMethod() int {
	t.parse()
	if t.parsed != nil {
		return t.parsed.NumMethod()
	}
	return 0
}

func (t *typeLazy) Child(i int) Type {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Child(i)
	}
	return nil
}

func (t *typeLazy) ChildByName(name string) (Type, bool) {
	t.parse()
	if t.parsed != nil {
		return t.parsed.ChildByName(name)
	}
	return nil, false
}

func (t *typeLazy) NumChild() int {
	t.parse()
	if t.parsed != nil {
		return t.parsed.NumChild()
	}
	return 0
}

func (t *typeLazy) Origin() ast.Node {
	return t.decl
}

func (t *typeLazy) Doc() *ast.CommentGroup {
	return t.comments
}

func (t *typeLazy) Comment() *ast.CommentGroup {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Comment()
	}
	return nil
}

func (t *typeLazy) NumTypeParam() int {
	t.parse()
	if t.parsed != nil {
		return t.parsed.NumTypeParam()
	}
	return 0
}

func (t *typeLazy) TypeParam(i int) Type {
	t.parse()
	if t.parsed != nil {
		return t.parsed.TypeParam(i)
	}
	return nil
}

func (t *typeLazy) Constraint() Type {
	t.parse()
	if t.parsed != nil {
		return t.parsed.Constraint()
	}
	return nil
}
