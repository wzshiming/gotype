package gotype

import (
	"go/ast"
	"go/token"
	"strings"
)

type importer interface {
	appendError(err error)
	Import(path string, src string) (Type, error)
	ImportName(path string, src string) (name string, goroot bool)
}

// typeName 解析表达式获取类型名字以及是否是导入的
func typeName(x ast.Expr) (name string, imported bool) {
	switch t := x.(type) {
	case *ast.Ident: // 当前包定义的
		return t.Name, false
	case *ast.SelectorExpr: // 外部导入的
		if _, ok := t.X.(*ast.Ident); ok {
			return t.Sel.Name, true
		}
	case *ast.StarExpr:
		return typeName(t.X)
	}
	return
}

func init() {
	for i := predeclaredTypesBeg + 1; i != predeclaredTypesEnd; i++ {
		k := strings.ToLower(i.String())
		predeclaredTypes[k] = i
	}
}

var predeclaredTypes = map[string]Kind{}

var tokenTypes = map[token.Token]Kind{
	token.INT:    Int,
	token.FLOAT:  Float64,
	token.IMAG:   Complex128,
	token.CHAR:   Int32,
	token.STRING: String,
}

var builtinFunc = map[string]builtinfunc{
	"append":  builtinfuncItem,
	"close":   builtinfuncVoid,
	"delete":  builtinfuncVoid,
	"panic":   builtinfuncVoid,
	"recover": builtinfuncInterface,
	"imag":    builtinfuncInt,
	"real":    builtinfuncInt,
	"make":    builtinfuncItem,
	"new":     builtinfuncPtrItem,
	"cap":     builtinfuncInt,
	"copy":    builtinfuncInt,
	"len":     builtinfuncInt,
}
