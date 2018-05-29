package gotype

import (
	"go/ast"
	"reflect"
	"strconv"
)

func (r *parser) EvalType(expr ast.Expr) Type {
	switch t := expr.(type) {
	case *ast.BadExpr:
		return nil
	case *ast.Ident:
		if k := predeclaredTypes[t.Name]; k != 0 {
			s := newTypeBuiltin(k, "")
			return s
		}
		s := newTypeNamed(t.Name, nil, r)
		return s
	case *ast.BasicLit:
		if k := tokenTypes[t.Kind]; k != 0 {
			s := newTypeBuiltin(k, t.Value)
			return s
		}
		return nil
	case *ast.FuncLit:
		return r.EvalType(t.Type)
	case *ast.CompositeLit:
		typ := r.EvalType(t.Type)
		if typ.Kind() != Struct {
			return typ
		}

		pairs := &typeValuePairs{}
		for _, v := range t.Elts {
			d := r.EvalType(v)
			pairs.li.Add(d)
		}
		return newTypeValueBind(typ, pairs)
	case *ast.ParenExpr:
		return r.EvalType(t.X)
	case *ast.SelectorExpr:
		s := r.EvalType(t.X)
		name := t.Sel.Name
		return newSelector(s, name)
	case *ast.IndexExpr:
		return r.EvalType(t.X).Elem()
	case *ast.SliceExpr:
		return r.EvalType(t.X)
	case *ast.TypeAssertExpr:
		return r.EvalType(t.Type)
	case *ast.CallExpr:
		switch b := t.Fun.(type) {
		case *ast.Ident:
			if bf, ok := builtinFunc[b.Name]; ok {
				switch bf {
				case builtinfuncInt:
					return newTypeBuiltin(Int, "")
				case builtinfuncPtrItem:
					return newTypePtr(r.EvalType(t.Args[0]))
				case builtinfuncItem:
					return r.EvalType(t.Args[0])
				case builtinfuncInterface:
					return newTypeBuiltin(Interface, "")
				case builtinfuncVoid:
					return newTypeBuiltin(Invalid, "")
				}
			}
		}

		b := r.EvalType(t.Fun)
		for b.Kind() == Var {
			b = b.Elem()
		}
		if b.Kind() == Func {
			l := b.NumOut()
			ts := make(types, 0, l)
			for i := 0; i != l; i++ {
				ts = append(ts, b.Out(i))
			}
			return newTypeTuple(ts)
		}
		return b
	case *ast.StarExpr:
		return newTypePtr(r.EvalType(t.X))
	case *ast.UnaryExpr:
		return r.EvalType(t.X)
	case *ast.BinaryExpr:
		return r.EvalType(t.X)
	case *ast.KeyValueExpr:
		k := r.EvalType(t.Key)
		v := r.EvalType(t.Value)
		return newTypeValuePair(k, v)
	case *ast.ArrayType:
		if t.Len == nil {
			return newTypeSlice(r.EvalType(t.Elt))
		} else {
			le := constValue(t.Len)
			i, _ := strconv.ParseInt(le, 0, 0)
			return newTypeArray(r.EvalType(t.Elt), int(i))
		}
	case *ast.StructType:
		s := &typeStruct{}

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
				t := &typeStructField{
					name: ty.Name(),
					elem: ty,
					tag:  tag,
				}
				s.anonymo.Add(t)
				continue
			}
			for _, name := range v.Names {
				t := &typeStructField{
					name: name.Name,
					elem: ty,
					tag:  tag,
				}
				s.fields.Add(t)
			}
		}
		return s
	case *ast.FuncType:
		s := &typeFunc{}
		if t.Params != nil {
			for _, v := range t.Params.List {
				if _, ok := v.Type.(*ast.Ellipsis); ok {
					s.variadic = true
				}
				ty := r.EvalType(v.Type)
				if ty == nil {
					continue
				}

				if v.Names == nil {
					t := newTypeVar("_", ty)
					s.params = append(s.params, t)
					continue
				}
				for _, name := range v.Names {
					t := newTypeVar(name.Name, ty)
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

				if v.Names == nil {
					t := newTypeVar("_", ty)
					s.results = append(s.results, t)
					continue
				}
				for _, name := range v.Names {
					t := newTypeVar(name.Name, ty)
					s.results = append(s.results, t)
				}
			}
		}
		return s
	case *ast.InterfaceType:
		s := &typeInterface{}
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
				t := newTypeAlias(name.Name, ty)
				s.methods.Add(t)
			}
		}
		return s
	case *ast.MapType:
		k := r.EvalType(t.Key)
		v := r.EvalType(t.Value)
		s := newTypeMap(k, v)
		return s
	case *ast.ChanType:
		v := r.EvalType(t.Value)
		s := newTypeChan(v, ChanDir(t.Dir))
		return s
	case *ast.Ellipsis:
		v := r.EvalType(t.Elt)
		s := newTypeSlice(v)
		return s
	default:
	}
	return nil
}
