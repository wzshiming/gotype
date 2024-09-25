package gotype

import (
	"go/ast"
	"go/token"
	"strings"
	"fmt"
)

type importer interface {
	appendError(err error)
	Import(path string, src string) (Type, error)
	ImportName(path string, src string) (name string, goroot bool)
}

// typeName Parses the expression to get the type name and whether it is imported
func typeName(x ast.Expr) (name string, imported bool) {
	switch t := x.(type) {
	case *ast.Ident: // Defined by the current package
		return t.Name, false
	case *ast.SelectorExpr: // Defined by the imported
		if _, ok := t.X.(*ast.Ident); ok {
			return t.Sel.Name, true
		}
	case *ast.StarExpr:
		return typeName(t.X)
	case *ast.IndexExpr:
		name, imported := typeName(t.X)
		if !imported {
			return name, false
		}
		param, imported := typeName(t.Index)
		if !imported {
			return name, false
		}
		return fmt.Sprintf("%s[%s]", name, param), true
	case *ast.IndexListExpr:
		name, imported := typeName(t.X)
		if !imported {
			return name, false
		}

		var params = make([]string, 0, len(t.Indices))
		for _, index := range t.Indices {
			param, imported := typeName(index)
			if !imported {
				return name, false
			}
			params = append(params, param)
		}
		return fmt.Sprintf("%s[%s]", name, strings.Join(params, ", ")), true
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
