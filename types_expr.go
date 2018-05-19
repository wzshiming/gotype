package gotype

import (
	"go/ast"
	"reflect"
	"strings"
)

func (r *Parser) GetTypes(expr ast.Expr) Type {
	switch t := expr.(type) {
	case *ast.BadExpr:
		return nil
	case *ast.Ident:
		if k := predeclaredTypes[t.Name]; k != 0 {
			s := &TypeBuiltin{}
			s.kind = k
			return s
		}
		s := &TypeNamed{}
		s.name = t.Name
		//		s.parser = r
		return s
	case *ast.BasicLit:
		if k := predeclaredTypes[strings.ToLower(t.Kind.String())]; k != 0 {
			s := &TypeBuiltin{}
			s.kind = k
			return s
		}
		return nil
	case *ast.FuncLit:
		return r.GetTypes(t.Type)
		//	case *ast.CompositeLit:
	case *ast.ParenExpr:
		return r.GetTypes(t.X)
	case *ast.SelectorExpr:
		return &TypeNamed{
			Type: r.GetTypes(t.X),
			name: t.Sel.Name,
		}
	case *ast.IndexExpr:
		return r.GetTypes(t.X).Elem()
	case *ast.SliceExpr:
		return r.GetTypes(t.X)
		// case *ast.TypeAssertExpr:
	case *ast.CallExpr:
		return r.GetTypes(t.Fun)
	case *ast.StarExpr:
		return r.GetTypes(t.X).Elem()
	case *ast.UnaryExpr:
		return r.GetTypes(t.X)
	case *ast.BinaryExpr:
		return r.GetTypes(t.X)
	// case *ast.KeyValueExpr:

	case *ast.ArrayType:
	// TODO
	case *ast.StructType:
		s := &TypeStruct{}
		s.parser = r
		s.kind = Struct
		if t.Fields == nil {
			return s
		}
		for _, v := range t.Fields.List {
			ty := r.GetTypes(v.Type)
			tag := reflect.StructTag(v.Tag.Value)
			if ty == nil {
				continue
			}
			for _, name := range v.Names {
				if name.Name == "" || name.Name == "_" {
					continue
				}
				s.fields = append(s.fields, &TypeStructField{
					Name: name.Name,
					Type: ty,
					Tag:  tag,
				})
			}
		}
		return s
	case *ast.FuncType:
		s := &TypeFunc{}
		if t.Params != nil {
			for _, v := range t.Params.List {
				ty := r.GetTypes(v.Type)
				if ty == nil {
					continue
				}
				for range v.Names {
					s.params = append(s.params, ty)
				}
			}
		}
		if t.Results != nil {
			for _, v := range t.Results.List {
				ty := r.GetTypes(v.Type)
				if ty == nil {
					continue
				}
				for range v.Names {
					s.results = append(s.results, ty)
				}
			}
		}
		return s
	case *ast.InterfaceType:
		s := &TypeInterface{}
		s.kind = Interface
		if t.Methods == nil {
			return s
		}
		for _, v := range t.Methods.List {
			ty := r.GetTypes(v.Type)
			if ty == nil {
				continue
			}
			for _, name := range v.Names {
				if name.Name == "" || name.Name == "_" {
					continue
				}
				s.methods = append(s.methods, &TypeMethod{
					Name: name.Name,
					Func: ty,
				})
			}
		}
		return s
	case *ast.MapType:
		// TODO
	case *ast.ChanType:
		// TODO
	}
	return nil
}
