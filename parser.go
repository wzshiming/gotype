package gotype

import (
	"go/ast"
	"go/token"
	"strconv"

	ffmt "gopkg.in/ffmt.v1"
)

type Parser struct {
	importer *Importer
	DotImp   []*Parser      // 尝试用点导入的包
	Imports  map[string]int // 导入的包

	Values map[string]Type          // 变量
	Funcs  map[string]Type          // 函数
	Method map[string][]*TypeMethod // 方法
	Types  map[string]Type          // 类型
}

// NewParser
func NewParser(i *Importer) *Parser {
	r := &Parser{
		importer: i,
		Imports:  map[string]int{},

		Values: map[string]Type{},
		Funcs:  map[string]Type{},
		Method: map[string][]*TypeMethod{},
		Types:  map[string]Type{},
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
		f := r.GetTypes(d.Type)
		if d.Recv != nil {
			name, ok := typeName(d.Recv.List[0].Type)
			if ok { // 不是当前包的方法
				return
			}

			r.Method[name] = append(r.Method[name], &TypeMethod{
				Name: d.Name.Name,
				Func: f,
			})
			return
		}
		r.Funcs[d.Name.Name] = f
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
				r.Imports[name] = 1
				if s.Name != nil && s.Name.Name == "." && r.importer != nil {
					p, err := r.importer.Import(name)
					if err != nil {
						ffmt.Mark(err)
						continue
					}
					r.DotImp = append(r.DotImp, p)
				}
			}
		case token.TYPE:
			for _, spec := range d.Specs {
				s, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				tt := r.GetTypes(s.Type)
				r.Types[s.Name.Name] = tt
			}
		}
	}
}

// ParserValue 解析
func (r *Parser) ParserValue(decl *ast.GenDecl) {
	var prev, name Type
	for _, spec := range decl.Specs {
		prev = name
		s, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}
		name = nil
		if s.Type != nil { // 有类型声明
			name = r.GetTypes(s.Type)
		} else if len(s.Values) == 0 { // 没有类型声明 但是一个常量  使用之前的类型
			if decl.Tok == token.CONST {
				name = prev
			}
		} else {
			// TODO: 还需要考虑多种情况
			name = r.GetTypes(s.Values[0])
		}
		for _, v := range s.Names {
			if v.Name == "" || v.Name == "_" {
				continue
			}
			r.Values[v.Name] = name
		}
	}
}
