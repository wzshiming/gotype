package gotype

import (
	"testing"
)

func TestImplements(t *testing.T) {

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

	scopeTime := Import(t, "time")
	val, ok := scopeTime.ChildByName("Time")
	if !ok {
		t.Fail()
		return
	}

	if !Identical(typ.Declaration().Declaration(), val) {
		t.Fail()
	}

	scopeEncodin := Import(t, "encoding")
	inter, ok := scopeEncodin.ChildByName("TextMarshaler")
	if !ok {
		t.Fail()
		return
	}

	if !Implements(val, inter) {
		t.Fail()
	}
}

func TestIdentical(t *testing.T) {
	var src = `package a
import time "time"

type A struct {
	String string
	Int int
	Array [8]byte
	Channel chan int
	Interface interface{}
	Map map[string]string
	Time time.Time
}
type B struct {
	String string
	Int int
	Array [8]byte
	Channel chan int
	Interface interface{}
	Map map[string]string
	Time time.Time
}
`

	scopeSrc := Parse(t, src)
	a, _ := scopeSrc.ChildByName("A")
	b, _ := scopeSrc.ChildByName("B")
	if !Identical(a, b) {
		t.Fail()
		return
	}
}

func TestEqual(t *testing.T) {
	scopeTime := Import(t, "time")

	val, ok := scopeTime.ChildByName("Time")
	if !ok {
		t.Fail()
		return
	}

	now, ok := scopeTime.ChildByName("Now")
	if !ok {
		t.Fail()
		return
	}

	if !Equal(val, now.Declaration().Out(0).Declaration()) {
		t.Fail()
	}
}
