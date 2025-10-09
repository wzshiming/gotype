package gotype

import (
	"runtime"
	"testing"
)

func BenchmarkEagerParsing(b *testing.B) {
	testdata := generateLargeSource(100) // 100 types
	
	b.ReportAllocs()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		imp := NewImporter()
		_, err := imp.ImportSource("_", []byte(testdata))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLazyParsing(b *testing.B) {
	testdata := generateLargeSource(100) // 100 types
	
	b.ReportAllocs()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		imp := NewImporter(WithLazyParsing())
		_, err := imp.ImportSource("_", []byte(testdata))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLazyParsingWithAccess(b *testing.B) {
	testdata := generateLargeSource(100) // 100 types
	
	b.ReportAllocs()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		imp := NewImporter(WithLazyParsing())
		scope, err := imp.ImportSource("_", []byte(testdata))
		if err != nil {
			b.Fatal(err)
		}
		
		// Access one type to trigger parsing
		if child, ok := scope.ChildByName("Type0"); ok {
			_ = child.Kind()
		}
	}
}

func TestMemoryUsageLazyVsEager(t *testing.T) {
	testdata := generateLargeSource(1000) // 1000 types for more visible difference
	
	// Measure eager parsing
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	
	impEager := NewImporter()
	scopeEager, err := impEager.ImportSource("_", []byte(testdata))
	if err != nil {
		t.Fatal(err)
	}
	
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)
	eagerAlloc := m2.Alloc - m1.Alloc
	
	// Keep scopeEager alive
	_ = scopeEager.NumChild()
	
	// Measure lazy parsing (without accessing)
	runtime.GC()
	var m3 runtime.MemStats
	runtime.ReadMemStats(&m3)
	
	impLazy := NewImporter(WithLazyParsing())
	scopeLazy, err := impLazy.ImportSource("_", []byte(testdata))
	if err != nil {
		t.Fatal(err)
	}
	
	runtime.GC()
	var m4 runtime.MemStats
	runtime.ReadMemStats(&m4)
	lazyAlloc := m4.Alloc - m3.Alloc
	
	// Keep scopeLazy alive
	_ = scopeLazy.NumChild()
	
	t.Logf("Eager parsing allocated: %d bytes", eagerAlloc)
	t.Logf("Lazy parsing allocated: %d bytes", lazyAlloc)
	
	if lazyAlloc < eagerAlloc {
		reduction := float64(eagerAlloc-lazyAlloc) / float64(eagerAlloc) * 100
		t.Logf("Lazy parsing reduced memory by %.2f%%", reduction)
	} else {
		t.Logf("Note: Memory measurements can vary; lazy parsing defers allocation until access")
	}
}

// generateLargeSource generates a Go source file with many types
func generateLargeSource(numTypes int) string {
	src := "package test\n\n"
	
	// Generate constants
	for i := 0; i < numTypes; i++ {
		src += "const Const" + string(rune('0'+i%10)) + string(rune('0'+(i/10)%10)) + string(rune('0'+(i/100)%10)) + " = " + string(rune('0'+i%10)) + "\n"
	}
	
	src += "\n"
	
	// Generate variables
	for i := 0; i < numTypes; i++ {
		src += "var Var" + string(rune('0'+i%10)) + string(rune('0'+(i/10)%10)) + string(rune('0'+(i/100)%10)) + " = " + string(rune('0'+i%10)) + "\n"
	}
	
	src += "\n"
	
	// Generate types
	for i := 0; i < numTypes; i++ {
		src += "type Type" + string(rune('0'+i%10)) + string(rune('0'+(i/10)%10)) + string(rune('0'+(i/100)%10)) + " struct {\n"
		src += "	Field1 string\n"
		src += "	Field2 int\n"
		src += "	Field3 bool\n"
		src += "}\n\n"
	}
	
	return src
}
