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
func (r *parser) ParserPackage(pkg *ast.Package) Type {
	for _, file := range pkg.Files {
		r.parserFile(file)
	}
	tt := newTypeScope(pkg.Name, r)
	tt = newTypeOrigin(tt, pkg, nil, nil)
	return tt
}

// ParserFile parser file
func (r *parser) parserFile(src *ast.File) {
	for _, decl := range src.Decls {
		r.parserDecl(decl)
	}
}

// ParserDecl parser declaration
func (r *parser) parserDecl(decl ast.Decl) {
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
	doc := decl.Doc
	f := r.EvalType(decl.Type)
	if decl.Recv != nil {
		name, ok := typeName(decl.Recv.List[0].Type)
		if ok {
			return
		}

		t := newTypeAlias(decl.Name.Name, f)
		t = newTypeOrigin(t, decl, doc, nil)
		b := r.method[name]
		b.Add(t)
		r.method[name] = b
		return
	}

	t := newTypeAlias(decl.Name.Name, f)
	t = newTypeOrigin(t, decl, doc, nil)
	r.nameds.Add(t)
	return
}

func (r *parser) parserType(decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		s, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		doc := s.Doc
		if decl.Lparen == 0 {
			doc = decl.Doc
		}
		comment := s.Comment

		tt := r.EvalType(s.Type)
		if s.Assign == 0 && tt.Kind() != Interface {
			tt = newTypeNamed(s.Name.Name, tt, r)
		} else {
			tt = newTypeAlias(s.Name.Name, tt)
		}

		tt = newTypeOrigin(tt, s, doc, comment)
		r.nameds.Add(tt)
	}
}

func (r *parser) parserImport(decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		s, ok := spec.(*ast.ImportSpec)
		if !ok {
			continue
		}

		doc := s.Doc
		if decl.Lparen == 0 {
			doc = decl.Doc
		}
		comment := s.Comment

		path, err := strconv.Unquote(s.Path.Value)
		if err != nil {
			continue
		}

		if r.importer == nil {
			continue
		}

		if s.Name == nil {
			tt := newTypeImport("", path, r.src, r.importer)
			tt = newTypeOrigin(tt, s, doc, comment)
			r.nameds.AddNoRepeat(tt)
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
					tt := p.Child(i)
					tt = newTypeOrigin(tt, s, doc, comment)
					r.nameds.AddNoRepeat(tt)
				}
			default:
				tt := newTypeImport(s.Name.Name, path, r.src, r.importer)
				tt = newTypeOrigin(tt, s, doc, comment)
				r.nameds.AddNoRepeat(tt)
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
		doc := s.Doc
		if decl.Lparen == 0 {
			doc = decl.Doc
		}
		comment := s.Comment

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
					tt := newTypeVar(v.Name, tup.all.Index(i))
					tt = newTypeOrigin(tt, s, doc, comment)
					r.nameds.Add(tt)
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
				tt := newTypeVar(v.Name, val)
				tt = newTypeOrigin(tt, s, doc, comment)
				r.nameds.Add(tt)
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
					tt := newTypeVar(v.Name, val)
					tt = newTypeOrigin(tt, s, doc, comment)
					r.nameds.Add(tt)
				} else {
					// No action
				}
			} else {
				if val == nil {
					tt := newTypeVar(v.Name, typ)
					tt = newTypeOrigin(tt, s, doc, comment)
					r.nameds.Add(tt)
				} else {
					tt := newTypeVar(v.Name, typ)
					tt = newTypeOrigin(tt, s, doc, comment)
					tt = newTypeValueBind(tt, val)
					r.nameds.Add(tt)
				}
			}

		}
	}
}
