package gotype

import "testing"

func TestImplements(t *testing.T) {
	scopeTime := Import(t, "time")

	scopeEncodin := Import(t, "encoding")

	val, ok := scopeTime.ChildByName("Time")
	if !ok {
		t.Fail()
		return
	}

	inter, ok := scopeEncodin.ChildByName("TextMarshaler")
	if !ok {
		t.Fail()
		return
	}

	if !Implements(val, inter) {
		t.Fail()
	}
}
