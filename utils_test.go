package gotype

import (
	"testing"
)

func Parse(t *testing.T, src string) Type {
	typ, err := NewImporter(WithCommentLocator()).ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return typ
}

func Import(t *testing.T, path string) Type {
	typ, err := NewImporter(WithCommentLocator()).Import(path)
	if err != nil {
		t.Fatal(err)
	}
	return typ
}
