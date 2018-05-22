package gotype

import (
	"go/ast"
	"reflect"
	"strconv"
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
		s := r.EvalType(t.X)
		name := t.Sel.Name
		if s.Kind() == Scope {
			return s.ChildByName(name)
		}
		b := s.MethodsByName(name)
		if b != nil {
			return b
		}
		b = s.FieldByName(name)
		if b != nil {
			return b
		}
		return nil
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
			le := constValue(t.Len)
			i, _ := strconv.ParseInt(le, 0, 0)
			return NewTypeArray(r.EvalType(t.Elt), int(i))
		}
	case *ast.StructType:
		s := &TypeStruct{}

		if t.Fields == nil {
			return s
		}
		for _, v := range t.Fields.List {
			ty := r.EvalType(v.Type)
			var tag reflect.StructTag
			if v.Tag != nil {
				tag = reflect.StructTag(v.Tag.Value)
			}
			if ty == nil {
				continue
			}

			if v.Names == nil {
				t := &TypeStructField{
					name: ty.Name(),
					typ:  ty,
					tag:  tag,
				}
				s.anonymo.Add(t)
				continue
			}
			for _, name := range v.Names {
				t := &TypeStructField{
					name: name.Name,
					typ:  ty,
					tag:  tag,
				}
				s.fields.Add(t)
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
				for _, name := range v.Names {
					t := NewTypeVar(name.Name, ty)
					s.params = append(s.params, t)
				}
			}
		}
		if t.Results != nil {
			for _, v := range t.Results.List {
				ty := r.EvalType(v.Type)
				if ty == nil {
					continue
				}
				for _, name := range v.Names {
					t := NewTypeVar(name.Name, ty)
					s.results = append(s.results, t)
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

			if v.Names == nil {
				s.anonymo.Add(ty)
			}

			for _, name := range v.Names {
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
