package gotype

import (
	"testing"
)

func Parse(t *testing.T, src string) Type {
	imp := getImporter(t)
	typ, err := imp.ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return typ
}

func Import(t *testing.T, path string) Type {
	imp := getImporter(t)
	typ, err := imp.Import(path, "")
	if err != nil {
		t.Fatal(err)
	}
	return typ

}

func getImporter(t *testing.T) *Importer {
	imp := NewImporter(
		WithCommentLocator(),
		ErrorHandler(func(err error) {
			t.Error(err)
		}),
	)
	return imp
}
