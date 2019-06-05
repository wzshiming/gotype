package gotype

func newTypeImport(name, path string, src string, imp importer) Type {
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
	imp   importer
	scope Type
}

func (t *typeImport) check() {
	if t.scope != nil {
		return
	}

	s, err := t.imp.Import(t.path, t.src)
	if err != nil {
		t.imp.appendError(err)
	}
	t.scope = s
}

func (t *typeImport) PkgPath() string {
	return t.path
}

func (t *typeImport) IsGoroot() bool {
	_, ok := t.imp.ImportName(t.path, t.src)
	return ok
}

func (t *typeImport) String() string {
	return t.path
}

func (t *typeImport) Name() string {
	if t.name != "" {
		return t.name
	}

	t.name, _ = t.imp.ImportName(t.path, t.src)
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
