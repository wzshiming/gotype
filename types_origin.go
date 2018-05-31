package gotype

import (
	"go/ast"
)

func newTypeOrigin(v Type, ori ast.Node, doc, comment *ast.CommentGroup) Type {
	return &typeOrigin{
		Type:    v,
		ori:     ori,
		doc:     doc,
		comment: comment,
	}
}

type typeOrigin struct {
	Type
	ori     ast.Node
	doc     *ast.CommentGroup
	comment *ast.CommentGroup
}

func (t *typeOrigin) Origin() ast.Node {
	return t.ori
}

func (t *typeOrigin) Doc() *ast.CommentGroup {
	return t.doc
}

func (t *typeOrigin) Comment() *ast.CommentGroup {
	return t.comment
}
