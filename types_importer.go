package gotype

import ffmt "gopkg.in/ffmt.v1"

func NewTypeImport(name, path string, imp *Importer) Type {
	return &TypeImport{
		name: name,
		path: path,
		imp:  imp,
	}
}

type TypeImport struct {
	typeBase
	name  string
	path  string
	imp   *Importer
	scope Type
}

func (t *TypeImport) check() {
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

func (t *TypeImport) Name() string {
	if t.name == "" {
		name, err := t.imp.ImportName(t.path)
		if err != nil {
			ffmt.Mark(err)
			return ""
		}
		t.name = name
	}
	return t.name
}

func (t *TypeImport) Kind() Kind {
	return Scope
}

func (t *TypeImport) ChildByName(name string) Type {
	t.check()
	if t.scope == nil {
		return nil
	}
	return t.scope.ChildByName(name)
}

func (t *TypeImport) Child(i int) Type {
	t.check()
	if t.scope == nil {
		return nil
	}
	return t.scope.Child(i)
}

func (t *TypeImport) NumChild() int {
	t.check()
	if t.scope == nil {
		return 0
	}
	return t.scope.NumChild()
}
