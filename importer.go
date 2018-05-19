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

func (i *Importer) Import(path string) (*Parser, error) {
	np := NewParser(i)
	p, err := parser.ParseDir(i.fset, path, i.filter, i.mode)
	if err != nil {
		return nil, err
	}
	for _, v := range p {
		np.ParserPackage(v)
	}

	return np, nil
}
