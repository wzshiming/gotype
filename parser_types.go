package gotype

import (
	"go/ast"
	"reflect"
	"strings"
)

func (r *Parser) EvalType(expr ast.Expr) Type {
	switch t := expr.(type) {
	case *ast.BadExpr:
		return nil
	case *ast.Ident:
		if k := predeclaredTypes[t.Name]; k != 0 {
			s := NewTypeBuiltin(k)
			return s
		}

		s := NewTypeNamed(t.Name, nil, r)
		return s
	case *ast.BasicLit:
		if k := predeclaredTypes[strings.ToLower(t.Kind.String())]; k != 0 {
			s := NewTypeBuiltin(k)
			return s
		}
		return nil
	case *ast.FuncLit:
		return r.EvalType(t.Type)
	case *ast.CompositeLit:
		return r.EvalType(t.Type)
	case *ast.ParenExpr:
		return r.EvalType(t.X)
	case *ast.SelectorExpr:
	// TODO:
	// a = b.c
	case *ast.IndexExpr:
		return r.EvalType(t.X).Elem()
	case *ast.SliceExpr:
		return r.EvalType(t.X)
		//	case *ast.TypeAssertExpr:

	case *ast.CallExpr:
		return r.EvalType(t.Fun)
	case *ast.StarExpr:
		return NewTypePtr(r.EvalType(t.X))
	case *ast.UnaryExpr:
		return r.EvalType(t.X)
	case *ast.BinaryExpr:
		return r.EvalType(t.X)
	// case *ast.KeyValueExpr:

	case *ast.ArrayType:
		if t.Len == nil {
			return NewTypeSlice(r.EvalType(t.Elt))
		} else {
			// TODO
			return NewTypeArray(r.EvalType(t.Elt), 1)
		}
	case *ast.StructType:
		s := &TypeStruct{}

		if t.Fields == nil {
			return s
		}
		for _, v := range t.Fields.List {
			ty := r.EvalType(v.Type)
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
				ty := r.EvalType(v.Type)
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
				ty := r.EvalType(v.Type)
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
		if t.Methods == nil {
			return s
		}
		for _, v := range t.Methods.List {
			ty := r.EvalType(v.Type)
			if ty == nil {
				continue
			}
			for _, name := range v.Names {
				if name.Name == "" || name.Name == "_" {
					continue
				}

				t := NewTypeNamed(name.Name, ty, r)
				s.methods.Add(t)
			}
		}
		return s
	case *ast.MapType:
		k := r.EvalType(t.Key)
		v := r.EvalType(t.Value)
		s := NewTypeMap(k, v)
		return s
	case *ast.ChanType:
		v := r.EvalType(t.Value)
		s := NewTypeChan(v, ChanDir(t.Dir))
		return s
	}
	return nil
}
