package gotype

import (
	"go/build"
	"go/parser"
	"go/token"
	"os"
)

type Importer struct {
	fset         *token.FileSet
	mode         parser.Mode
	bufType      map[string]Type
	bufBuild     map[string]*build.Package
	errorHandler func(error)
}

func NewImporter(options ...option) *Importer {
	i := &Importer{
		fset:     token.NewFileSet(),
		mode:     parser.ParseComments,
		bufType:  map[string]Type{},
		bufBuild: map[string]*build.Package{},
		errorHandler: func(err error) {
			return
		},
	}
	for _, v := range options {
		v(i)
	}
	return i
}

func (i *Importer) importBuild(path string, src string) (*build.Package, bool) {
	k := path + " " + src
	if v, ok := i.bufBuild[k]; ok {
		return v, true
	}

	imp, err := build.Import(path, src, 0)
	if err != nil {
		i.errorHandler(err)
		return nil, false
	}
	i.bufBuild[k] = imp
	return imp, true
}

func (i *Importer) importName(path string, src string) (name string, goroot bool) {
	imp, ok := i.importBuild(path, src)
	if !ok {
		return "", false
	}
	return imp.Name, imp.Goroot
}

func (i *Importer) Import(path string) Type {
	return i.impor(path, ".")
}

func (i *Importer) impor(path string, src string) Type {

	imp, ok := i.importBuild(path, src)
	if !ok {
		return nil
	}
	dir := imp.Dir

	tt, ok := i.bufType[dir]
	if ok {
		return tt
	}

	m := map[string]bool{}
	for _, v := range imp.GoFiles {
		m[v] = true
	}

	p, err := parser.ParseDir(i.fset, dir, func(fi os.FileInfo) bool {
		return m[fi.Name()]
	}, i.mode)

	if err != nil {
		i.errorHandler(err)
		return nil
	}

	for name, v := range p {
		np := newParser(i, dir)
		np.ParserPackage(v)
		t := newTypeScope(name, np)
		i.bufType[dir] = t
		return t
	}
	return nil
}
