package gotype

import (
	"go/ast"
	"strings"
)

func constValue(x ast.Expr) string {
	switch t := x.(type) {
	case *ast.BasicLit:
		return t.Value
	}
	return ""
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
