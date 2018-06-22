package gotype

import (
	"go/ast"
	"go/token"
	"sort"
)

func newTypeCommentLocatorComment(typ Type, pos, end token.Pos, comments []*ast.CommentGroup) Type {
	return &typeCommentLocatorComment{
		Type:     typ,
		pos:      pos,
		end:      end,
		comments: comments,
	}
}

type typeCommentLocatorComment struct {
	Type
	pos      token.Pos
	end      token.Pos
	comments []*ast.CommentGroup
	comment  *ast.CommentGroup
}

func (t *typeCommentLocatorComment) Comment() *ast.CommentGroup {
	if t.comment == nil {
		t.comment = commentLocator(t.pos, t.end, t.comments)
	}
	return t.comment
}

func commentLocator(pos, end token.Pos, comments []*ast.CommentGroup) *ast.CommentGroup {
	i := sort.Search(len(comments), func(i int) bool {
		return pos < comments[i].Pos()
	})
	if i == -1 {
		return nil
	}
	if comments[i].End() < end {
		return comments[i]
	}
	return nil
}
