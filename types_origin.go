package gotype

import (
	"go/ast"
)

func newTypeOrigin(v Type, ori ast.Node, info *info, doc, comment *ast.CommentGroup) Type {
	return &typeOrigin{
		Type:    v,
		info:    info,
		ori:     ori,
		doc:     doc,
		comment: comment,
	}
}

type typeOrigin struct {
	Type
	ori     ast.Node
	info    *info
	goroot  bool
	doc     *ast.CommentGroup
	comment *ast.CommentGroup
	pkgPath string
}

func (t *typeOrigin) Origin() ast.Node {
	return t.ori
}

func (t *typeOrigin) PkgPath() string {
	if t.pkgPath != "" {
		return t.pkgPath
	}
	t.pkgPath = t.Type.PkgPath()
	if t.pkgPath == "" {
		t.pkgPath = t.info.PkgPath
	}
	return t.pkgPath
}

func (t *typeOrigin) IsGoroot() bool {
	return t.info.Goroot || t.Type.IsGoroot()
}

func (t *typeOrigin) Doc() *ast.CommentGroup {
	if t.doc != nil {
		return t.doc
	}
	return t.Type.Doc()
}

func (t *typeOrigin) Comment() *ast.CommentGroup {
	if t.comment != nil {
		return t.comment
	}
	return t.Type.Comment()
}
