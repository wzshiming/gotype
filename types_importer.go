package gotype

func newTypeImport(name, path string, src string, imp *Importer) Type {
	return &typeImport{
		name: name,
		path: path,
		src:  src,
		imp:  imp,
	}
}

type typeImport struct {
	typeBase
	name  string
	path  string
	src   string
	imp   *Importer
	scope Type
}

func (t *typeImport) check() {
	if t.scope != nil {
		return
	}

	s, err := t.imp.impor(t.path, t.src)
	if err != nil {
		t.imp.errorHandler(err)
	}
	t.scope = s
}

func (t *typeImport) String() string {
	return t.name
}

func (t *typeImport) Name() string {
	if t.name == "" {
		name, _ := t.imp.importName(t.path, t.src)
		t.name = name
	}
	return t.name
}

func (t *typeImport) Kind() Kind {
	return Scope
}

func (t *typeImport) ChildByName(name string) (Type, bool) {
	t.check()
	if t.scope == nil {
		return nil, false
	}
	return t.scope.ChildByName(name)
}

func (t *typeImport) Child(i int) Type {
	t.check()
	if t.scope == nil {
		return nil
	}
	return t.scope.Child(i)
}

func (t *typeImport) NumChild() int {
	t.check()
	if t.scope == nil {
		return 0
	}
	return t.scope.NumChild()
}
