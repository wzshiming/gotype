package gotype

import (
	"go/ast"
	"go/token"
	"strconv"
)

type parser struct {
	importer         importer
	isCommentLocator bool
	info             *info
	comments         []*ast.CommentGroup
}

// NewParser
func newParser(i importer, c bool, src string, pkg string, goroot bool) *parser {
	r := &parser{
		importer:         i,
		isCommentLocator: c,
		info:             newInfo(src, pkg, goroot),
	}
	return r
}

// ParsePackage parse package
func (r *parser) ParsePackage(pkg *ast.Package) Type {
	for filename, file := range pkg.Files {
		r.parseFile(r.info.File(filename), file)
	}
	tt := newTypeScope(pkg.Name, r.info)
	tt = newTypeOrigin(tt, pkg, r.info, nil, nil)
	return tt
}

// ParseFile parse file
func (r *parser) ParseFile(file *ast.File) Type {
	r.parseFile(r.info.File(""), file)
	tt := newTypeScope(file.Name.String(), r.info)
	tt = newTypeOrigin(tt, file, r.info, nil, nil)
	return tt
}

// parseFile parse file
func (r *parser) parseFile(info *infoFile, file *ast.File) {
	r.comments = file.Comments
	for _, decl := range file.Decls {
		r.parseDecl(info, decl)
	}
	r.comments = nil
}

// parseDecl parse declaration
func (r *parser) parseDecl(info *infoFile, decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		r.parseFunc(info, d)
	case *ast.GenDecl:
		switch d.Tok {
		case token.CONST, token.VAR:
			r.parseValue(info, d)
		case token.IMPORT:
			r.parseImport(info, d)
		case token.TYPE:
			r.parseType(info, d)
		}
	}
}

// parseFunc parse function
func (r *parser) parseFunc(info *infoFile, decl *ast.FuncDecl) {
	doc := decl.Doc
	t := r.evalType(info, decl.Type)
	t = newDeclaration(decl.Name.Name, t)
	t = newTypeOrigin(t, decl, r.info, doc, nil)

	if decl.Recv != nil {
		name, ok := typeName(decl.Recv.List[0].Type)
		if ok {
			return
		}
		info.AddMethod(name, t)
		return
	}
	info.AddType(t)
	return
}

// parseType parse type
func (r *parser) parseType(info *infoFile, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		s, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		doc := decl.Doc
		if s.Doc != nil {
			doc = s.Doc
		}
		comment := s.Comment

		tt := r.evalType(info, s.Type)
		if s.Assign == 0 && tt.Kind() != Interface {
			tt = newTypeNamed(s.Name.Name, tt, info)
		} else {
			tt = newTypeAlias(s.Name.Name, tt)
		}

		tt = newTypeOrigin(tt, s, r.info, doc, comment)
		info.AddType(tt)
	}
}

// parserImport parser import
func (r *parser) parseImport(info *infoFile, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		s, ok := spec.(*ast.ImportSpec)
		if !ok {
			continue
		}

		doc := decl.Doc
		if s.Doc != nil {
			doc = s.Doc
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
			tt := newTypeImport("", path, r.info.Src, r.importer)
			tt = newTypeOrigin(tt, s, r.info, doc, comment)
			info.AddPkg(tt)
		} else {
			switch s.Name.Name {
			case "_":
			case ".":
				p, err := r.importer.Import(path, r.info.PkgPath)
				if err != nil {
					r.importer.appendError(err)
					continue
				}
				l := p.NumChild()
				for i := 0; i != l; i++ {
					tt := p.Child(i)
					tt = newTypeOrigin(tt, s, r.info, doc, comment)
					info.AddPkg(tt)
				}
			default:
				tt := newTypeImport(s.Name.Name, path, r.info.Src, r.importer)
				tt = newTypeOrigin(tt, s, r.info, doc, comment)
				info.AddPkg(tt)
			}
		}
	}
}

// parseValue parse value
func (r *parser) parseValue(info *infoFile, decl *ast.GenDecl) {
	var prevVal ast.Expr
	var prevTyp ast.Expr

loop:
	for index, spec := range decl.Specs {
		s, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}

		doc := decl.Doc
		if s.Doc != nil {
			doc = s.Doc
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
			typ = r.evalType(info, ctyp)
			if decl.Tok == token.CONST {
				prevTyp = ctyp
			}
		}

		var val Type

		switch l := len(values); l {
		case 0:
		case 1:
			val = r.evalType(info, values[0])
			if decl.Tok == token.CONST {
				prevVal = values[0]
				val = newEvalBind(val, int64(index), info)
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
					info.AddType(tt)
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
				val = r.evalType(info, values[i])
				if decl.Tok == token.CONST {
					val = newEvalBind(val, int64(index), info)
				}
				tt := newDeclaration(v.Name, val)
				tt = newTypeOrigin(tt, s, r.info, doc, comment)
				info.AddType(tt)
			}
			continue loop
		}

		for _, v := range s.Names {
			if v.Name == "" {
				continue
			}

			if typ == nil {
				if val != nil {
					tt := newDeclaration(v.Name, val)
					tt = newTypeOrigin(tt, s, r.info, doc, comment)
					info.AddType(tt)
				} else {
					// No action
				}
			} else {
				if val == nil {
					tt := newDeclaration(v.Name, typ)
					tt = newTypeOrigin(tt, s, r.info, doc, comment)
					info.AddType(tt)
				} else {
					tt := newDeclaration(v.Name, typ)
					tt = newTypeOrigin(tt, s, r.info, doc, comment)
					tt = newTypeValueBind(tt, val, r.info)
					info.AddType(tt)
				}
			}
		}
	}
}
