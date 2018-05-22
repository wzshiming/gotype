package gotype

import ffmt "gopkg.in/ffmt.v1"

func newTypeImport(name, path string, imp *Importer) Type {
	return &typeImport{
		name: name,
		path: path,
		imp:  imp,
	}
}

type typeImport struct {
	typeBase
	name  string
	path  string
	imp   *Importer
	scope Type
}

func (t *typeImport) check() {
	if t.scope != nil {
		return
	}

	s, err := t.imp.Import(t.path)
	if err != nil {
		ffmt.Mark(err)
		return
	}

	t.scope = s
}

func (t *typeImport) Name() string {
	if t.name == "" {
		name, _, err := t.imp.ImportName(t.path)
		if err != nil {
			ffmt.Mark(err)
			return ""
		}
		t.name = name
	}
	return t.name
}

func (t *typeImport) Kind() Kind {
	return Scope
}

func (t *typeImport) ChildByName(name string) Type {
	t.check()
	if t.scope == nil {
		return nil
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
