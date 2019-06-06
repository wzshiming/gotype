package gotype

import (
	"fmt"
	"go/ast"
	"go/build"
	goparser "go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// Importer Go source type analyzer
type Importer struct {
	fset             *token.FileSet
	mode             goparser.Mode
	bufType          map[string]Type
	bufBuild         map[string]*build.Package
	errorHandler     func(error)
	importHandler    func(path, src, dir string)
	isCommentLocator bool
	ctx              build.Context
}

// NewImporter creates a new importer
func NewImporter(options ...Option) *Importer {
	i := &Importer{
		fset:     token.NewFileSet(),
		mode:     goparser.ParseComments,
		bufType:  map[string]Type{},
		bufBuild: map[string]*build.Package{},
		errorHandler: func(err error) {
			return
		},
		ctx: build.Default,
	}
	for _, v := range options {
		v(i)
	}
	return i
}

// ImportPackage returns go package scope
func (i *Importer) ImportPackage(path string, pkg *ast.Package) (Type, error) {
	t, ok := i.bufType[path]
	if ok {
		return t, nil
	}
	np := newParser(i, i.isCommentLocator, path, false)
	t = np.ParsePackage(pkg)
	i.bufType[path] = t
	return t, nil
}

// ImportFile returns go package scope
func (i *Importer) ImportFile(path string, f *ast.File) (Type, error) {
	t, ok := i.bufType[path]
	if ok {
		return t, nil
	}
	np := newParser(i, i.isCommentLocator, path, false)
	t = np.ParseFile(f)
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

// FileSet returns the FileSet
func (i *Importer) FileSet() *token.FileSet {
	return i.fset
}

// ImportBuild returns details about the Go package named by the import path.
func (i *Importer) ImportBuild(path string, src string) (*build.Package, error) {
	dot := filepath.HasPrefix(src, ".")
	src = filepath.Clean(src)
	gopath := filepath.Join(i.ctx.GOPATH, "src")
	rsrc := src

	if dot {
		pwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		rsrc = filepath.Join(pwd, src)
	} else if filepath.HasPrefix(src, "/") {
		rsrc = src
	} else if !filepath.HasPrefix(src, gopath) {
		rsrc = filepath.Join(gopath, src)
	}

	k := path + " " + rsrc
	if v, ok := i.bufBuild[k]; ok {
		return v, nil
	}
	imp, err := i.ctx.Import(path, rsrc, 0)
	if err != nil {
		i.appendError(err)
		return nil, err
	}
	i.bufBuild[k] = imp
	return imp, nil
}

// appendError append error
func (i *Importer) appendError(err error) {
	if i.errorHandler != nil {
		i.errorHandler(err)
	}
}

// ImportName returns go package name
func (i *Importer) ImportName(path string, src string) (name string, goroot bool) {
	imp, err := i.ImportBuild(path, src)
	if err != nil {
		return "", false
	}
	return imp.Name, imp.Goroot
}

// Import returns go package scope
func (i *Importer) Import(path string, src string) (Type, error) {
	imp, err := i.ImportBuild(path, src)
	if err != nil {
		return nil, err
	}
	dir := imp.Dir

	tt, ok := i.bufType[dir]
	if ok {
		return tt, nil
	}
	if i.importHandler != nil {
		i.importHandler(path, src, dir)
	}

	m := map[string]bool{}
	for _, v := range imp.GoFiles {
		m[v] = true
	}

	p, err := goparser.ParseDir(i.fset, dir, func(fi os.FileInfo) bool {
		return m[fi.Name()]
	}, i.mode)

	if err != nil {
		i.appendError(err)
		return nil, err
	}

	for _, v := range p {
		np := newParser(i, i.isCommentLocator, imp.ImportPath, imp.Goroot)
		t := np.ParsePackage(v)
		i.bufType[dir] = t
		return t, nil
	}
	err = fmt.Errorf(`No go source code was found under the package path "%s"`, path)
	i.appendError(err)
	return nil, err
}
