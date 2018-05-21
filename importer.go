package gotype

import (
	"go/build"
	"go/parser"
	"go/token"
	"os"
)

type Importer struct {
	fset *token.FileSet
	mode parser.Mode
	buf  map[string]Type
}

func NewImporter() *Importer {
	return &Importer{
		fset: token.NewFileSet(),
		mode: parser.ParseComments,
		buf:  map[string]Type{},
	}
}

func (i *Importer) ImportName(path string) (string, error) {
	imp, err := build.Import(path, ".", build.FindOnly)
	if err != nil {
		return "", err
	}

	return imp.Name, nil
}

func (i *Importer) Import(path string) (Type, error) {
	if v, ok := i.buf[path]; ok {
		return v, nil
	}

	imp, err := build.Import(path, ".", 0)
	if err != nil {
		return nil, err
	}

	m := map[string]bool{}
	for _, v := range imp.GoFiles {
		m[v] = true
	}

	dir := imp.Dir
	p, err := parser.ParseDir(i.fset, dir, func(fi os.FileInfo) bool {
		return m[fi.Name()]
	}, i.mode)

	if err != nil {
		return nil, err
	}

	for name, v := range p {
		np := NewParser(i)
		np.ParserPackage(v)
		t := NewTypeScope(name, np)
		i.buf[path] = t
		return t, nil
	}
	return nil, nil
}
