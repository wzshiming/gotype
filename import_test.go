package gotype

import "testing"

func TestImport(t *testing.T) {
	const src = `
package a

import "time"
var now = time.Now()
`
	scope := Parse(t, src)

	typ, ok := scope.ChildByName("now")
	if !ok {
		t.Fail()
	}

	if typ.Name() != "now" {
		t.Fail()
	}
	if typ.Kind() != Declaration {
		t.Fail()
	}
	typ = typ.Declaration()

	if typ.Name() != "_" {
		t.Fail()
	}
	if typ.Kind() != Declaration {
		t.Fail()
	}
	typ = typ.Declaration()

	if typ.Name() != "Time" {
		t.Fail()
	}
	if typ.PkgPath() != "time" {
		t.Fail()
	}
	if typ.Doc() == nil {
		t.Fail()
	}
}
