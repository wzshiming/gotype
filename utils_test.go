package gotype

import (
	"testing"
)

func Parse(t *testing.T, src string) Type {
	typ, err := NewImporter(
		WithCommentLocator(),
		ErrorHandler(func(err error) {
			t.Error(err)
		}),
	).ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return typ
}

func Import(t *testing.T, path string) Type {
	typ, err := NewImporter(
		WithCommentLocator(),
		ErrorHandler(func(err error) {
			t.Error(err)
		}),
	).Import(path)
	if err != nil {
		t.Fatal(err)
	}
	return typ
}
