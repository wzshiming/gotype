package gotype

import (
	"go/ast"
	"go/token"
	"strconv"
)

type parser struct {
	importer *Importer
	method   map[string]types // type method
	nameds   types            // var, func, type, packgae
	src      string
}

// NewParser
func newParser(i *Importer, src string) *parser {
	r := &parser{
		importer: i,
		method:   map[string]types{},
		nameds:   types{},
		src:      src,
	}
	return r
}

// ParserFile parser package
func (r *parser) ParserPackage(pkg *ast.Package) {
	for _, file := range pkg.Files {
		r.ParserFile(file)
	}
}

// ParserFile parser file
func (r *parser) ParserFile(src *ast.File) {
	for _, decl := range src.Decls {
		r.ParserDecl(decl)
	}
}

// ParserDecl parser declaration
func (r *parser) ParserDecl(decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		r.ParserFunc(d)
	case *ast.GenDecl:
		switch d.Tok {
		case token.CONST, token.VAR:
			r.parserValue(d)
		case token.IMPORT:
			r.parserImport(d)
		case token.TYPE:
			r.parserType(d)
		}
	}
}

func (r *parser) ParserFunc(decl *ast.FuncDecl) {
	f := r.EvalType(decl.Type)
	if decl.Recv != nil {
		name, ok := typeName(decl.Recv.List[0].Type)
		if ok {
			return
		}

		t := newTypeAlias(decl.Name.Name, f)
		b := r.method[name]
		b.Add(t)
		r.method[name] = b
		return
	}

	t := newTypeAlias(decl.Name.Name, f)
	r.nameds.Add(t)
	return
}

func (r *parser) parserType(decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		s, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		tt := r.EvalType(s.Type)
		if s.Assign == 0 && tt.Kind() != Interface {
			tt = newTypeNamed(s.Name.Name, tt, r)
		} else {
			tt = newTypeAlias(s.Name.Name, tt)
		}
		r.nameds.Add(tt)
	}
}

func (r *parser) parserImport(decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		s, ok := spec.(*ast.ImportSpec)
		if !ok {
			continue
		}
		path, err := strconv.Unquote(s.Path.Value)
		if err != nil {
			continue
		}

		if r.importer == nil {
			continue
		}

		if s.Name == nil {
			p := newTypeImport("", path, r.src, r.importer)
			r.nameds.AddNoRepeat(p)
		} else {
			switch s.Name.Name {
			case "_":
			case ".":
				p := r.importer.impor(path, r.src)
				if p == nil {
					continue
				}
				l := p.NumChild()
				for i := 0; i != l; i++ {
					r.nameds.AddNoRepeat(p.Child(i))
				}
			default:
				t := newTypeImport(s.Name.Name, path, r.src, r.importer)
				r.nameds.AddNoRepeat(t)
			}
		}
	}
}

func (r *parser) parserValue(decl *ast.GenDecl) {
	var prev, val Type
loop:
	for _, spec := range decl.Specs {
		prev = val
		s, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}

		var typ Type
		if s.Type != nil { // Type definition
			typ = r.EvalType(s.Type)
		}
		val = nil

		switch l := len(s.Values); l {
		case 0:
			if decl.Tok == token.CONST {
				val = prev
			}
		case 1:
			val = r.EvalType(s.Values[0])
			if tup, ok := val.(*typeTuple); ok {

				l := tup.all.Len()
				for i, v := range s.Names {
					if v.Name == "" || v.Name == "_" {
						continue
					}
					if i == l {
						break
					}
					t := newTypeVar(v.Name, tup.all.Index(i))
					r.nameds.Add(t)
				}
				continue loop
			}
		default:
			l := len(s.Values)
			for i, v := range s.Names {
				if v.Name == "" || v.Name == "_" {
					continue
				}
				if i == l {
					break
				}
				val := r.EvalType(s.Values[i])
				t := newTypeVar(v.Name, val)
				r.nameds.Add(t)
			}
			continue loop
		}

		if val == nil {
			continue
		}
		for _, v := range s.Names {
			if v.Name == "" || v.Name == "_" {
				continue
			}

			if typ == nil {
				if val != nil {
					t := newTypeVar(v.Name, val)
					r.nameds.Add(t)
				} else {
					// No action
				}
			} else {
				if val == nil {
					t := newTypeVar(v.Name, typ)
					r.nameds.Add(t)
				} else {
					t := newTypeVar(v.Name, typ)
					t = newTypeValueBind(t, val)
					r.nameds.Add(t)
				}
			}

		}
	}
}
