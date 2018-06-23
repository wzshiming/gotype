package gotype

import (
	"fmt"
	"go/ast"
	"go/build"
	goparser "go/parser"
	"go/token"
	"os"
)

// Importer Go source type analyzer
type Importer struct {
	fset             *token.FileSet
	mode             goparser.Mode
	bufType          map[string]Type
	bufBuild         map[string]*build.Package
	errorHandler     func(error)
	isCommentLocator bool
}

// NewImporter creates a new importer
func NewImporter(options ...option) *Importer {
	i := &Importer{
		fset:     token.NewFileSet(),
		mode:     goparser.ParseComments,
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

// Import returns go package scope
func (i *Importer) Import(path string) (Type, error) {
	return i.importParse(path, ".")
}

// ImportPackage returns go package scope
func (i *Importer) ImportPackage(path string, pkg *ast.Package) (Type, error) {
	np := newParser(i.importParseErrorHandler, i.isCommentLocator, path, false)
	t := np.ParsePackage(pkg)
	i.bufType[path] = t
	return t, nil
}

// ImportFile returns go package scope
func (i *Importer) ImportFile(path string, f *ast.File) (Type, error) {
	np := newParser(i.importParseErrorHandler, i.isCommentLocator, path, false)
	t := np.ParseFile(f)
	i.bufType[path] = t
	return t, nil
}

// ImportSource returns go package scope
func (i *Importer) ImportSource(path string, src []byte) (Type, error) {
	f, err := goparser.ParseFile(i.fset, path, src, i.mode)
	if err != nil {
		return nil, err
	}
	return i.ImportFile(path, f)
}

// ImportBuild returns details about the Go package named by the import path.
func (i *Importer) ImportBuild(path string) (*build.Package, error) {
	imp, err := i.importBuild(path, ".")
	if err != nil {
		return nil, err
	}
	return imp, nil
}

// FileSet returns the FileSet
func (i *Importer) FileSet() *token.FileSet {
	return i.fset
}

func (i *Importer) importBuild(path string, src string) (*build.Package, error) {
	k := path + " " + src
	if v, ok := i.bufBuild[k]; ok {
		return v, nil
	}

	imp, err := build.Import(path, src, 0)
	if err != nil {
		return nil, err
	}
	i.bufBuild[k] = imp
	return imp, nil
}

func (i *Importer) importName(path string, src string) (name string, goroot bool) {
	imp, err := i.importBuild(path, src)
	if err != nil {
		return "", false
	}
	return imp.Name, imp.Goroot
}

func (i *Importer) importParseErrorHandler(path string, src string) (Type, error) {
	t, err := i.importParse(path, src)
	if err != nil {
		i.errorHandler(err)
	}
	return t, err
}

func (i *Importer) importParse(path string, src string) (Type, error) {
	imp, err := i.importBuild(path, src)
	if err != nil {
		return nil, err
	}
	dir := imp.Dir

	tt, ok := i.bufType[dir]
	if ok {
		return tt, nil
	}

	m := map[string]bool{}
	for _, v := range imp.GoFiles {
		m[v] = true
	}

	p, err := goparser.ParseDir(i.fset, dir, func(fi os.FileInfo) bool {
		return m[fi.Name()]
	}, i.mode)

	if err != nil {
		return nil, err
	}

	for _, v := range p {
		np := newParser(i.importParseErrorHandler, i.isCommentLocator, imp.ImportPath, imp.Goroot)
		t := np.ParsePackage(v)
		i.bufType[dir] = t
		return t, nil
	}
	return nil, fmt.Errorf(`No go source code was found under the package path "%s"`, path)
}
