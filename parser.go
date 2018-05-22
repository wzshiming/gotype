package gotype

import (
	"go/ast"
	"go/token"
	"strconv"
)

type Parser struct {
	importer *Importer
	DotImp   Types            // 尝试用点导入的包
	Imports  Types            // 导入的包
	Method   map[string]Types // 方法
	Nameds   Types            // 变量 函数 类型
}

// NewParser
func NewParser(i *Importer) *Parser {
	r := &Parser{
		importer: i,
		Imports:  Types{},
		Method:   map[string]Types{},
		Nameds:   Types{},
	}
	return r
}

// ParserFile 解析文件
func (r *Parser) ParserFile(src *ast.File) {
	// 解析全部顶级关键字
	for _, decl := range src.Decls {
		r.ParserDecl(decl)
	}
}

func (r *Parser) ParserPackage(pkg *ast.Package) {
	for _, file := range pkg.Files {
		r.ParserFile(file)
	}
}

// ParserDecl 解析顶级关键字
func (r *Parser) ParserDecl(decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		f := r.EvalType(d.Type)
		if d.Recv != nil {
			name, ok := typeName(d.Recv.List[0].Type)
			if ok { // 不是当前包的方法
				return
			}

			t := NewTypeNamed(d.Name.Name, f, r)
			b := r.Method[name]
			b.Add(t)
			r.Method[name] = b
			return
		}

		t := NewTypeNamed(d.Name.Name, f, r)
		r.Nameds.Add(t)
		return
	case *ast.GenDecl:
		switch d.Tok {
		case token.CONST, token.VAR:
			r.ParserValue(d)
		case token.IMPORT:
			for _, spec := range d.Specs {
				s, ok := spec.(*ast.ImportSpec)
				if !ok {
					continue
				}
				name, err := strconv.Unquote(s.Path.Value)
				if err != nil {
					continue
				}

				if r.importer == nil {
					continue
				}

				if s.Name == nil {
					p := NewTypeImport("", name, r.importer)
					r.Imports.Add(p)
				} else {
					switch s.Name.Name {
					case "_":
					case ".":
					// TODO
					//r.DotImp = append(r.DotImp, p)
					default:

						t := NewTypeImport(s.Name.Name, name, r.importer)
						r.Imports.Add(t)
					}
				}
			}
		case token.TYPE:
			for _, spec := range d.Specs {
				s, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				tt := r.EvalType(s.Type)
				if s.Assign == 0 {
					tt = NewTypeNamed(s.Name.Name, tt, r)
				} else {
					tt = NewTypeAlias(s.Name.Name, tt)
				}
				r.Nameds.Add(tt)
			}
		}
	}
}

// ParserValue 解析
func (r *Parser) ParserValue(decl *ast.GenDecl) {
	var prev, val Type
	for _, spec := range decl.Specs {
		prev = val
		s, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}
		val = nil
		if s.Type != nil { // 有类型声明
			val = r.EvalType(s.Type)
		} else if len(s.Values) == 0 { // 没有类型声明 但是一个常量  使用之前的类型
			if decl.Tok == token.CONST {
				val = prev
			}
		} else {
			// TODO: 还需要考虑多种情况

			val = r.EvalType(s.Values[0])
			//ffmt.P(s.Values)
		}
		for _, v := range s.Names {
			if v.Name == "" || v.Name == "_" {
				continue
			}

			t := NewTypeVar(v.Name, val)
			r.Nameds.Add(t)
		}
	}
}
