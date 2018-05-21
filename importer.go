package gotype

import (
	"go/parser"
	"go/token"
	"os"
)

type Importer struct {
	fset   *token.FileSet
	mode   parser.Mode
	filter func(os.FileInfo) bool
}

func NewImporter() *Importer {

	return &Importer{
		fset:   token.NewFileSet(),
		mode:   parser.ParseComments,
		filter: func(os.FileInfo) bool { return true },
	}
}

func (i *Importer) Import(path string) (Types, error) {

	p, err := parser.ParseDir(i.fset, path, i.filter, i.mode)
	if err != nil {
		return nil, err
	}

	pkg := make(Types, 0, len(p))
	for name, v := range p {
		np := NewParser(i)
		np.ParserPackage(v)
		t := NewTypeScope(name, np)
		pkg = append(pkg, t)
	}

	return pkg, nil
}
