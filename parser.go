package gotype

import (
	"go/ast"
	"go/token"
	"strconv"
)

type parser struct {
	importer         importParseFunc
	isCommentLocator bool
	info             *info
	comments         []*ast.CommentGroup
}

// NewParser
func newParser(i importParseFunc, c bool, pkg string, goroot bool) *parser {
	r := &parser{
		importer:         i,
		isCommentLocator: c,
		info:             newInfo(pkg, goroot),
	}
	return r
}

// ParsePackage parse package
func (r *parser) ParsePackage(pkg *ast.Package) Type {
	for _, file := range pkg.Files {
		r.parseFile(file)
	}
	tt := newTypeScope(pkg.Name, r.info)
	tt = newTypeOrigin(tt, pkg, r.info, nil, nil)
	return tt
}

// ParseFile parse file
func (r *parser) ParseFile(file *ast.File) Type {
	r.parseFile(file)
	tt := newTypeScope(file.Name.String(), r.info)
	tt = newTypeOrigin(tt, file, r.info, nil, nil)
	return tt
}

// parseFile parse file
func (r *parser) parseFile(file *ast.File) {
	r.comments = file.Comments
	for _, decl := range file.Decls {
		r.parseDecl(decl)
	}
	r.comments = nil
}

// parseDecl parse declaration
func (r *parser) parseDecl(decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		r.parseFunc(d)
	case *ast.GenDecl:
		switch d.Tok {
		case token.CONST, token.VAR:
			r.parseValue(d)
		case token.IMPORT:
			r.parseImport(d)
		case token.TYPE:
			r.parseType(d)
		}
	}
}

// parseFunc parse function
func (r *parser) parseFunc(decl *ast.FuncDecl) {
	doc := decl.Doc
	t := r.EvalType(decl.Type)
	t = newDeclaration(decl.Name.Name, t)
	t = newTypeOrigin(t, decl, r.info, doc, nil)

	if decl.Recv != nil {
		name, ok := typeName(decl.Recv.List[0].Type)
		if ok {
			return
		}
		b := r.info.Methods[name]
		b.Add(t)
		r.info.Methods[name] = b
		return
	}
	r.info.Named.Add(t)
	return
}

// parseType parse type
func (r *parser) parseType(decl *ast.GenDecl) {
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
			tt = newTypeNamed(s.Name.Name, tt, r.info)
		} else {
			tt = newTypeAlias(s.Name.Name, tt)
		}

		tt = newTypeOrigin(tt, s, r.info, doc, comment)
		r.info.Named.Add(tt)
	}
}

// parserImport parser import
func (r *parser) parseImport(decl *ast.GenDecl) {
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
			tt := newTypeImport("", path, r.info.PkgPath, r.importer)
			tt = newTypeOrigin(tt, s, r.info, doc, comment)
			r.info.Named.AddNoRepeat(tt)
		} else {
			switch s.Name.Name {
			case "_":
			case ".":
				p, err := r.importer(path, r.info.PkgPath)
				if err != nil {
					continue
				}
				l := p.NumChild()
				for i := 0; i != l; i++ {
					tt := p.Child(i)
					tt = newTypeOrigin(tt, s, r.info, doc, comment)
					r.info.Named.AddNoRepeat(tt)
				}
			default:
				tt := newTypeImport(s.Name.Name, path, r.info.PkgPath, r.importer)
				tt = newTypeOrigin(tt, s, r.info, doc, comment)
				r.info.Named.AddNoRepeat(tt)
			}
		}
	}
}

// parseValue parse value
func (r *parser) parseValue(decl *ast.GenDecl) {
	var prevVal ast.Expr
	var prevTyp ast.Expr

loop:
	for index, spec := range decl.Specs {
		s, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}
		doc := s.Doc
		if decl.Lparen == 0 {
			doc = decl.Doc
		}
		comment := s.Comment

		ctyp := s.Type
		values := s.Values

		if decl.Tok == token.CONST && s.Type == nil && len(values) == 0 {
			values = append(values, prevVal)
			ctyp = prevTyp
		}

		var typ Type
		if ctyp != nil { // Type definition
			typ = r.EvalType(ctyp)
			if decl.Tok == token.CONST {
				prevTyp = ctyp
			}
		}

		var val Type

		switch l := len(values); l {
		case 0:
			continue loop
		case 1:
			val = r.EvalType(values[0])
			if decl.Tok == token.CONST {
				prevVal = values[0]
				val = newEvalBind(val, int64(index), r.info)
			}
			if tup, ok := val.(*typeTuple); ok {
				l := tup.all.Len()
				for i, v := range s.Names {
					if v.Name == "" {
						continue
					}
					if i == l {
						break
					}
					val = tup.all.Index(i)
					tt := newDeclaration(v.Name, val)
					tt = newTypeOrigin(tt, s, r.info, doc, comment)
					r.info.Named.Add(tt)
				}
				continue loop
			}
		default:
			l := len(values)
			for i, v := range s.Names {
				if v.Name == "" {
					continue
				}
				if i == l {
					break
				}
				val = r.EvalType(values[i])
				if decl.Tok == token.CONST {
					val = newEvalBind(val, int64(index), r.info)
				}
				tt := newDeclaration(v.Name, val)
				tt = newTypeOrigin(tt, s, r.info, doc, comment)
				r.info.Named.Add(tt)
			}
			continue loop
		}

		if val == nil {
			continue
		}
		for _, v := range s.Names {
			if v.Name == "" {
				continue
			}

			if typ == nil {
				if val != nil {
					tt := newDeclaration(v.Name, val)
					tt = newTypeOrigin(tt, s, r.info, doc, comment)
					r.info.Named.Add(tt)
				} else {
					// No action
				}
			} else {
				if val == nil {
					tt := newDeclaration(v.Name, typ)
					tt = newTypeOrigin(tt, s, r.info, doc, comment)
					r.info.Named.Add(tt)
				} else {
					tt := newDeclaration(v.Name, typ)
					tt = newTypeOrigin(tt, s, r.info, doc, comment)
					tt = newTypeValueBind(tt, val, r.info)
					r.info.Named.Add(tt)
				}
			}
		}
	}
}
