package gotype

import (
	"strings"
	"testing"
)

func TestFuncParameterComment(t *testing.T) {
	var testdata = []string{
		`package a
func Func (A string /* A */,
B int, // B
) (
	C int, // C
) {}
`,
	}
	for _, src := range testdata {
		testFuncParameterComment(t, Parse(t, src))
	}
}

func testFuncParameterComment(t *testing.T, scope Type) {

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
	var testdata = []string{
		`package a

type Type struct {}

// Medhod1
func(Type) Medhod1(){}

// Medhod2
func(Type) Medhod2(string)string{}
`,
	}
	for _, src := range testdata {
		testFieldMedhodDoc(t, Parse(t, src))
	}
}

func testFieldMedhodDoc(t *testing.T, scope Type) {

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

func TestInterfaceMedhodComment(t *testing.T) {
	var testdata = []string{
		`package a

type Type interface {
	Medhod1() // Medhod1
	Medhod2(string) string // Medhod2
}
`,
	}
	for _, src := range testdata {
		testInterfaceMedhodComment(t, Parse(t, src))
	}
}

func testInterfaceMedhodComment(t *testing.T, scope Type) {

	typ, ok := scope.ChildByName("Type")
	if !ok {
		t.Fail()
	}

	num := typ.NumMethod()
	for i := 0; i != num; i++ {
		v := typ.Method(i)
		name := v.Name()
		doc := strings.TrimSpace(v.Comment().Text())
		if name != doc {
			t.Fatal("Error field comment:", name, doc)
		}
	}
}

func TestInterfaceMedhodDoc(t *testing.T) {
	var testdata = []string{
		`package a

type Type interface {
	// Medhod1
	Medhod1()
	// Medhod2
	Medhod2(string) string
}
`,
	}
	for _, src := range testdata {
		testInterfaceMedhodDoc(t, Parse(t, src))
	}
}

func testInterfaceMedhodDoc(t *testing.T, scope Type) {

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
			t.Fatal("Error field doc:", name, doc)
		}
	}
}

func TestFieldComment(t *testing.T) {
	var testdata = []string{
		`package a

type Type struct {
	Field1 string  // Field1
	
	Field2 struct{
	
	} // Field2
}
`,
	}
	for _, src := range testdata {
		testFieldComment(t, Parse(t, src))
	}
}

func testFieldComment(t *testing.T, scope Type) {

	typ, ok := scope.ChildByName("Type")
	if !ok {
		t.Fail()
	}

	num := typ.NumField()
	for i := 0; i != num; i++ {
		v := typ.Field(i)
		name := v.Name()
		doc := strings.TrimSpace(v.Comment().Text())
		if name != doc {
			t.Fatal("Error field comment:", name, doc)
		}
	}
}

func TestFieldDoc(t *testing.T) {
	var testdata = []string{
		`package a

type Type struct {
	// Field1
	Field1 string 
	
	// Field2
	Field2 struct{
	
	}
}
`,
	}
	for _, src := range testdata {
		testFieldDoc(t, Parse(t, src))
	}
}

func testFieldDoc(t *testing.T, scope Type) {

	typ, ok := scope.ChildByName("Type")
	if !ok {
		t.Fail()
	}

	num := typ.NumField()
	for i := 0; i != num; i++ {
		v := typ.Field(i)
		name := v.Name()
		doc := strings.TrimSpace(v.Doc().Text())
		if name != doc {
			t.Fatal("Error field doc:", name, doc)
		}
	}
}

func TestDoc(t *testing.T) {
	var testdata = []string{
		`package a

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

`,
	}
	for _, src := range testdata {
		testDoc(t, Parse(t, src))
	}
}

func testDoc(t *testing.T, scope Type) {

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
	var testdata = []string{
		`package a

const Const = "" // Const

var Var = 0 // Var

type Type struct {
	
} // Type

type Type2 struct {} // Type2

`,
	}
	for _, src := range testdata {
		testComment(t, Parse(t, src))
	}
}

func testComment(t *testing.T, scope Type) {

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
