package gotype

import (
	"fmt"
	"strings"
	"testing"
)

func TestImport(t *testing.T) {
	const src = `
package a

import "time"
var now = time.Now()
`
	scope, err := NewImporter().ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}
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

func TestKind(t *testing.T) {
	const src = `
package a
type Bool bool
type Int int
type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Uint uint
type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64
type Uintptr uintptr
type Float32 float32
type Float64 float64
type Complex64 complex64
type Complex128 complex128
type String string
type Byte byte
type Rune rune
type Error error
type Array [1]struct{}
type Chan chan struct{}
type Func func()
type Interface interface{}
type Map map[string]struct{}
type Ptr *int
type Slice []struct{}
`
	scope, err := NewImporter().ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}

	num := scope.NumChild()
	for i := 0; i != num; i++ {
		v := scope.Child(i)
		name := v.Name()
		kind := v.Kind().String()
		if name != kind {
			t.Fatal("Error kind:", name, kind)
		}
	}
}

func TestFuncParameterComment(t *testing.T) {
	const src = `
package a

func Func (A string /* A */,
B int, // B
) (
	C int, // C
) {}
`
	scope, err := NewImporter(WithCommentLocator()).ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}

	typ, ok := scope.ChildByName("Func")
	if !ok {
		t.Fail()
	}
	typ = typ.Declaration()

	{
		num := typ.NumIn()
		for i := 0; i != num; i++ {
			v := typ.In(i)
			name := v.Name()
			comment := strings.TrimSpace(v.Comment().Text())
			if name != comment {
				t.Fatal("Error func parameter comment:", name, comment)
			}
		}
	}

	{
		num := typ.NumOut()
		for i := 0; i != num; i++ {
			v := typ.Out(i)
			name := v.Name()
			comment := strings.TrimSpace(v.Comment().Text())
			if name != comment {
				t.Fatal("Error func parameter comment:", name, comment)
			}
		}
	}
}

func TestFieldMedhodDoc(t *testing.T) {
	const src = `
package a

type Type struct {}

// Medhod
func(Type) Medhod(){}

`
	scope, err := NewImporter().ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}

	typ, ok := scope.ChildByName("Type")
	if !ok {
		t.Fail()
	}

	num := typ.NumMethod()
	for i := 0; i != num; i++ {
		v := typ.Method(i)
		name := v.Name()
		doc := strings.TrimSpace(v.Doc().Text())
		if name != doc {
			t.Fatal("Error method doc:", name, doc)
		}
	}
}

func TestInterfaceMedhodDocAndComment(t *testing.T) {
	const src = `
package a

type Type interface {
	// Medhod doc
	Medhod() // Medhod comment
}
`
	scope, err := NewImporter().ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}

	typ, ok := scope.ChildByName("Type")
	if !ok {
		t.Fail()
	}

	num := typ.NumMethod()
	for i := 0; i != num; i++ {
		v := typ.Method(i)
		{
			name := v.Name() + " doc"
			doc := strings.TrimSpace(v.Doc().Text())
			if name != doc {
				t.Fatal("Error field doc:", name, doc)
			}
		}

		{
			name := v.Name() + " comment"
			comment := strings.TrimSpace(v.Comment().Text())
			if name != comment {
				t.Fatal("Error field comment:", name, comment)
			}
		}
	}
}

func TestFieldDocAndComment(t *testing.T) {
	const src = `
package a

type Type struct {
	// Field1 doc
	Field1 string // Field1 comment
	
	
	// Field1 doc
	Field1 struct{
	
	} // Field1 comment
}


`
	scope, err := NewImporter().ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}

	typ, ok := scope.ChildByName("Type")
	if !ok {
		t.Fail()
	}

	num := typ.NumField()
	for i := 0; i != num; i++ {
		v := typ.Field(i)
		{
			name := v.Name() + " doc"
			doc := strings.TrimSpace(v.Doc().Text())
			if name != doc {
				t.Fatal("Error field doc:", name, doc)
			}
		}

		{
			name := v.Name() + " comment"
			comment := strings.TrimSpace(v.Comment().Text())
			if name != comment {
				t.Fatal("Error field comment:", name, comment)
			}
		}
	}
}

func TestDoc(t *testing.T) {
	const src = `
package a

// Const
const Const = ""

var (
	// Var
	Var = 0
		
)

// Type
type Type struct {}

// Func
func Func() {
}

`
	scope, err := NewImporter().ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}

	num := scope.NumChild()
	for i := 0; i != num; i++ {
		v := scope.Child(i)
		name := v.Name()
		doc := strings.TrimSpace(v.Doc().Text())
		if name != doc {
			t.Fatal("Error doc:", name, doc)
		}
	}
}

func TestComment(t *testing.T) {
	const src = `
package a

const Const = "" // Const

var Var = 0 // Var

type Type struct {
	
} // Type

type Type2 struct {} // Type2

`
	scope, err := NewImporter().ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}

	num := scope.NumChild()
	for i := 0; i != num; i++ {
		v := scope.Child(i)
		name := v.Name()
		comment := strings.TrimSpace(v.Comment().Text())
		if name != comment {
			t.Fatal("Error comment:", name, comment)
		}
	}
}

func TestConstValue(t *testing.T) {
	const src = `
package a

// A
const a = "A"


const (
	// CAB
	b = c + "B"
	// CA
	c = "C" + a
)
`
	scope, err := NewImporter().ImportSource("_", []byte(src))
	if err != nil {
		t.Fatal(err)
	}

	num := scope.NumChild()
	for i := 0; i != num; i++ {
		v := scope.Child(i)
		val := v.Value()
		doc := fmt.Sprintf(`"%s"`, strings.TrimSpace(v.Doc().Text()))
		if val != doc {
			t.Fatal("Error value method:", val, doc)
		}
	}
}
