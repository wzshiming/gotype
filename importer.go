package gotype

import (
	"go/build"
	"go/parser"
	"go/token"
	"os"
)

type Importer struct {
	fset     *token.FileSet
	mode     parser.Mode
	bufType  map[string]Type
	bufBuild map[string]*build.Package
}

func NewImporter() *Importer {
	return &Importer{
		fset:     token.NewFileSet(),
		mode:     parser.ParseComments,
		bufType:  map[string]Type{},
		bufBuild: map[string]*build.Package{},
	}
}

func (i *Importer) importBuild(path string) (*build.Package, error) {
	if v, ok := i.bufBuild[path]; ok {
		return v, nil
	}
	imp, err := build.Import(path, ".", 0)
	if err != nil {
		return nil, err
	}
	i.bufBuild[path] = imp
	return imp, nil
}

func (i *Importer) ImportName(path string) (name string, goroot bool, err error) {
	imp, err := i.importBuild(path)
	if err != nil {
		return "", false, err
	}

	return imp.Name, imp.Goroot, nil
}

func (i *Importer) Import(path string) (Type, error) {
	if v, ok := i.bufType[path]; ok {
		return v, nil
	}

	imp, err := i.importBuild(path)
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
		np := newParser(i)
		np.ParserPackage(v)
		t := newTypeScope(name, np)
		i.bufType[path] = t
		return t, nil
	}
	return nil, nil
}
