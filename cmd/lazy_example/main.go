package main

import (
	"fmt"
	"runtime"

	"github.com/wzshiming/gotype"
)

func main() {
	// Example source code to parse
	sourceCode := `package example

// Person represents a person with basic information
type Person struct {
	Name string
	Age  int
}

// Company represents a company
type Company struct {
	Name    string
	Founded int
}

const (
	MaxAge = 100
	MinAge = 0
)

var (
	DefaultPerson = Person{Name: "Unknown", Age: 0}
)

// Greet returns a greeting message
func Greet(name string) string {
	return "Hello, " + name
}
`

	fmt.Println("=== Eager Parsing ===")
	demonstrateEagerParsing(sourceCode)

	fmt.Println("\n=== Lazy Parsing ===")
	demonstrateLazyParsing(sourceCode)
}

func demonstrateEagerParsing(sourceCode string) {
	// Measure memory before
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// Create importer with eager parsing (default)
	imp := gotype.NewImporter()
	scope, err := imp.ImportSource("example", []byte(sourceCode))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Measure memory after
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	fmt.Printf("Memory allocated: %d bytes\n", m2.Alloc-m1.Alloc)
	fmt.Printf("Number of types: %d\n", scope.NumChild())

	// Access a type
	if person, ok := scope.ChildByName("Person"); ok {
		fmt.Printf("Found type: %s (kind: %v)\n", person.Name(), person.Kind())
	}
}

func demonstrateLazyParsing(sourceCode string) {
	// Measure memory before
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// Create importer with lazy parsing enabled
	imp := gotype.NewImporter(gotype.WithLazyParsing())
	scope, err := imp.ImportSource("example", []byte(sourceCode))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Measure memory after initial parsing (before accessing types)
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	fmt.Printf("Memory allocated (before accessing types): %d bytes\n", m2.Alloc-m1.Alloc)
	fmt.Printf("Number of types: %d\n", scope.NumChild())

	// Now access a type (this triggers lazy parsing of that specific type)
	if person, ok := scope.ChildByName("Person"); ok {
		fmt.Printf("Found type: %s (kind: %v)\n", person.Name(), person.Kind())
		
		// Check if it's a struct and access fields
		if person.Kind() == gotype.Struct {
			fmt.Printf("Number of fields: %d\n", person.NumField())
		}
	}

	// Measure memory after accessing one type
	runtime.GC()
	var m3 runtime.MemStats
	runtime.ReadMemStats(&m3)
	fmt.Printf("Memory allocated (after accessing Person): %d bytes\n", m3.Alloc-m1.Alloc)
}
