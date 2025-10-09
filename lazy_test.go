package gotype

import (
	"testing"
)

func TestLazyParsing(t *testing.T) {
	testdata := `package a

// Const1 is a constant
const Const1 = "hello"

// Const2 is another constant
const Const2 = 42

var (
	// Var1 is a variable
	Var1 = 0
	// Var2 is another variable
	Var2 = "world"
)

// Type1 is a struct
type Type1 struct {
	Field1 string
	Field2 int
}

// Type2 is another struct
type Type2 struct {
	Field3 bool
}

// Func1 is a function
func Func1() {
}
`

	// Test with lazy parsing enabled
	impLazy := NewImporter(WithLazyParsing(), WithCommentLocator())
	scopeLazy, err := impLazy.ImportSource("_", []byte(testdata))
	if err != nil {
		t.Fatal(err)
	}

	// Test with lazy parsing disabled (default)
	impEager := NewImporter(WithCommentLocator())
	scopeEager, err := impEager.ImportSource("_", []byte(testdata))
	if err != nil {
		t.Fatal(err)
	}

	// Verify both produce the same number of children
	numLazy := scopeLazy.NumChild()
	numEager := scopeEager.NumChild()
	if numLazy != numEager {
		t.Fatalf("Different number of children: lazy=%d, eager=%d", numLazy, numEager)
	}

	// Verify we can access children by name
	testCases := []string{"Const1", "Const2", "Var1", "Var2", "Type1", "Type2", "Func1"}
	for _, name := range testCases {
		lazyType, okLazy := scopeLazy.ChildByName(name)
		eagerType, okEager := scopeEager.ChildByName(name)

		if okLazy != okEager {
			t.Fatalf("ChildByName(%s) returned different results: lazy=%v, eager=%v", name, okLazy, okEager)
		}

		if !okLazy {
			t.Fatalf("ChildByName(%s) not found", name)
		}

		// Verify names match
		if lazyType.Name() != eagerType.Name() {
			t.Fatalf("Names differ for %s: lazy=%s, eager=%s", name, lazyType.Name(), eagerType.Name())
		}

		// Verify kinds match
		if lazyType.Kind() != eagerType.Kind() {
			t.Fatalf("Kinds differ for %s: lazy=%v, eager=%v", name, lazyType.Kind(), eagerType.Kind())
		}

		// For struct types, verify fields
		if lazyType.Kind() == Declaration {
			lazyDecl := lazyType.Declaration()
			eagerDecl := eagerType.Declaration()

			if lazyDecl.Kind() != eagerDecl.Kind() {
				t.Fatalf("Declaration kinds differ for %s: lazy=%v, eager=%v", name, lazyDecl.Kind(), eagerDecl.Kind())
			}

			if lazyDecl.Kind() == Struct {
				if lazyDecl.NumField() != eagerDecl.NumField() {
					t.Fatalf("Field count differs for %s: lazy=%d, eager=%d", name, lazyDecl.NumField(), eagerDecl.NumField())
				}
			}
		}
	}
}

func TestLazyParsingConstants(t *testing.T) {
	testdata := `package a

const (
	C1 = 1
	C2 = 2
	C3 = "hello"
)
`

	impLazy := NewImporter(WithLazyParsing())
	scopeLazy, err := impLazy.ImportSource("_", []byte(testdata))
	if err != nil {
		t.Fatal(err)
	}

	// Access constants
	c1, ok := scopeLazy.ChildByName("C1")
	if !ok {
		t.Fatal("C1 not found")
	}

	if c1.Kind() != Declaration {
		t.Fatalf("C1 kind is %v, expected Declaration", c1.Kind())
	}

	c2, ok := scopeLazy.ChildByName("C2")
	if !ok {
		t.Fatal("C2 not found")
	}

	if c2.Kind() != Declaration {
		t.Fatalf("C2 kind is %v, expected Declaration", c2.Kind())
	}

	c3, ok := scopeLazy.ChildByName("C3")
	if !ok {
		t.Fatal("C3 not found")
	}

	if c3.Kind() != Declaration {
		t.Fatalf("C3 kind is %v, expected Declaration", c3.Kind())
	}
}

func TestLazyParsingIteration(t *testing.T) {
	testdata := `package a

const A = 1
const B = 2
var C = 3
type D struct{}
func E() {}
`

	impLazy := NewImporter(WithLazyParsing())
	scopeLazy, err := impLazy.ImportSource("_", []byte(testdata))
	if err != nil {
		t.Fatal(err)
	}

	// Iterate through all children
	num := scopeLazy.NumChild()
	if num != 5 {
		t.Fatalf("Expected 5 children, got %d", num)
	}

	expectedNames := map[string]bool{
		"A": false,
		"B": false,
		"C": false,
		"D": false,
		"E": false,
	}

	for i := 0; i < num; i++ {
		child := scopeLazy.Child(i)
		name := child.Name()
		if _, ok := expectedNames[name]; !ok {
			t.Fatalf("Unexpected child: %s", name)
		}
		expectedNames[name] = true
	}

	for name, found := range expectedNames {
		if !found {
			t.Fatalf("Child %s not found during iteration", name)
		}
	}
}
