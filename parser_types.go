package gotype

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

func (r *parser) EvalType(expr ast.Expr) (ret Type) {
	return r.evalType(r.info.File(""), expr)
}

func (r *parser) evalType(info *infoFile, expr ast.Expr) (ret Type) {
	defer func() {
		if ret != nil {
			ret = newTypeOrigin(ret, expr, r.info, nil, nil)
		}
	}()
	switch t := expr.(type) {
	case *ast.BadExpr:
		return nil
	case *ast.Ident:
		if k := predeclaredTypes[t.Name]; k != 0 {
			s := newTypeBuiltin(k, "")
			return s
		}
		s := newTypeNamed(t.Name, nil, info)
		return s
	case *ast.BasicLit:
		if k := tokenTypes[t.Kind]; k != 0 {
			s := newTypeBuiltin(k, t.Value)
			return s
		}
		return nil
	case *ast.FuncLit:
		return r.evalType(info, t.Type)
	case *ast.CompositeLit:
		typ := r.evalType(info, t.Type)
		if typ.Kind() != Struct {
			return typ
		}

		pairs := &typeValuePairs{}
		for _, v := range t.Elts {
			d := r.evalType(info, v)
			pairs.li.Add(d)
		}
		return newTypeValueBind(typ, pairs, r.info)
	case *ast.ParenExpr:
		return r.evalType(info, t.X)
	case *ast.SelectorExpr:
		s := r.evalType(info, t.X)
		name := t.Sel.Name
		return newSelector(s, name)
	case *ast.IndexExpr:
		return r.evalType(info, t.X).Elem()
	case *ast.SliceExpr:
		return r.evalType(info, t.X)
	case *ast.TypeAssertExpr:
		return r.evalType(info, t.Type)
	case *ast.CallExpr:
		switch b := t.Fun.(type) {
		case *ast.Ident:
			if bf, ok := builtinFunc[b.Name]; ok {
				switch bf {
				case builtinfuncInt:
					return newTypeBuiltin(Int, "")
				case builtinfuncPtrItem:
					return newTypePtr(r.evalType(info, t.Args[0]))
				case builtinfuncItem:
					return r.evalType(info, t.Args[0])
				case builtinfuncInterface:
					return newTypeBuiltin(Interface, "")
				case builtinfuncVoid:
					return newTypeBuiltin(Invalid, "")
				}
			}
		}

		b := r.evalType(info, t.Fun)
		for b.Kind() == Declaration {
			b = b.Declaration()
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
		return newTypePtr(r.evalType(info, t.X))
	case *ast.UnaryExpr:
		if t.Op == token.AND {
			return newTypePtr(r.evalType(info, t.X))
		}
		return r.evalType(info, t.X)
	case *ast.BinaryExpr:
		return r.evalType(info, t.X)
	case *ast.KeyValueExpr:
		k := r.evalType(info, t.Key)
		v := r.evalType(info, t.Value)
		return newTypeValuePair(k, v)
	case *ast.ArrayType:
		if t.Len == nil {
			return newTypeSlice(r.evalType(info, t.Elt))
		}

		if ell, ok := t.Len.(*ast.Ellipsis); ok && ell.Ellipsis != 0 {
			return newTypeArray(r.evalType(info, t.Elt), 0)
		}

		length := newEvalBind(r.evalType(info, t.Len), 0, info)
		i, _ := strconv.ParseInt(length.Value(), 0, 0)
		return newTypeArray(r.evalType(info, t.Elt), int(i))
	case *ast.StructType:
		s := &typeStruct{}

		if t.Fields == nil {
			return s
		}
		for _, v := range t.Fields.List {
			ty := r.evalType(info, v.Type)
			var tag reflect.StructTag
			if v.Tag != nil {
				tv := v.Tag.Value
				tv = strings.Trim(tv, "`")
				tag = reflect.StructTag(tv)
			}

			if ty == nil {
				continue
			}

			if v.Names == nil {
				t := &typeStructField{
					name:      ty.Name(),
					elem:      ty,
					tag:       tag,
					anonymous: true,
				}
				tt := newTypeOrigin(t, v, r.info, v.Doc, v.Comment)
				s.fields = append(s.fields, tt)
				continue
			}
			for _, name := range v.Names {
				n := name.Name
				if n == "" {
					n = "_"
				}
				t := &typeStructField{
					name: n,
					elem: ty,
					tag:  tag,
				}
				tt := newTypeOrigin(t, v, r.info, v.Doc, v.Comment)
				s.fields = append(s.fields, tt)
			}
		}
		return s
	case *ast.FuncType:
		s := &typeFunc{}
		if t.Params != nil {
			list := t.Params.List
			for pk, v := range list {
				if _, ok := v.Type.(*ast.Ellipsis); ok {
					s.variadic = true
				}
				ty := r.evalType(info, v.Type)
				if ty == nil {
					continue
				}

				if v.Names == nil {
					t := newDeclaration("_", ty)
					s.params = append(s.params, t)
					continue
				}
				for nk, name := range v.Names {
					tv := newDeclaration(name.Name, ty)
					tt := newTypeOrigin(tv, v, r.info, v.Doc, v.Comment)

					// Parameter comment, Because the default criteria are not handled.
					if r.isCommentLocator {
						var beg, end token.Pos
						if nk != len(v.Names)-1 {
							beg = name.End()
						} else {
							beg = v.Type.End()
						}

						if nk != len(v.Names)-1 {
							end = v.Names[nk+1].Pos()
						} else if pk != len(list)-1 {
							end = list[pk+1].Pos()
						} else {
							end = t.Params.End()
						}
						tt = newTypeCommentLocatorComment(tt, beg, end, r.comments)
					}

					s.params = append(s.params, tt)
				}
			}
		}
		if t.Results != nil {
			list := t.Results.List
			for pk, v := range list {
				ty := r.evalType(info, v.Type)
				if ty == nil {
					continue
				}

				if v.Names == nil {
					t := newDeclaration("_", ty)
					s.results = append(s.results, t)
					continue
				}
				for nk, name := range v.Names {
					tv := newDeclaration(name.Name, ty)
					tt := newTypeOrigin(tv, v, r.info, v.Doc, v.Comment)

					// Parameter comment, Because the default criteria are not handled.
					if r.isCommentLocator {
						var beg, end token.Pos
						if nk != len(v.Names)-1 {
							beg = name.End()
						} else {
							beg = v.Type.End()
						}
						if nk != len(v.Names)-1 {
							end = v.Names[nk+1].Pos()
						} else if pk != len(list)-1 {
							end = list[pk+1].Pos()
						} else {
							end = t.Results.End()
						}
						tt = newTypeCommentLocatorComment(tt, beg, end, r.comments)
					}

					s.results = append(s.results, tt)
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
			ty := r.evalType(info, v.Type)
			if ty == nil {
				continue
			}

			if v.Names == nil {
				t := newDeclaration(ty.Name(), ty)
				tt := newTypeOrigin(t, v, r.info, v.Doc, v.Comment)
				s.all.Add(tt)
				s.anonymo.Add(tt)
			} else {
				for _, name := range v.Names {
					t := newDeclaration(name.Name, ty)
					tt := newTypeOrigin(t, v, r.info, v.Doc, v.Comment)
					s.all.Add(tt)
					s.method.Add(tt)
				}
			}
		}
		return s
	case *ast.MapType:
		k := r.evalType(info, t.Key)
		v := r.evalType(info, t.Value)
		s := newTypeMap(k, v)
		return s
	case *ast.ChanType:
		v := r.evalType(info, t.Value)
		s := newTypeChan(v, ChanDir(t.Dir))
		return s
	case *ast.Ellipsis:
		v := r.evalType(info, t.Elt)
		s := newTypeSlice(v)
		return s
	default:
	}
	return nil
}
