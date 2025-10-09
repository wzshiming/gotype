# gotype

Golang source code parsing, usage like reflect package

[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/gotype)](https://goreportcard.com/report/github.com/wzshiming/gotype)
[![GoDoc](https://godoc.org/github.com/wzshiming/gotype?status.svg)](https://godoc.org/github.com/wzshiming/gotype)
[![GitHub license](https://img.shields.io/github/license/wzshiming/gotype.svg)](https://github.com/wzshiming/gotype/blob/master/LICENSE)

- [English](https://github.com/wzshiming/gotype/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/gotype/blob/master/README_cn.md)

## Usage

[API Documentation](https://godoc.org/github.com/wzshiming/gotype)

[Examples](https://github.com/wzshiming/gotype/blob/master/cmd/pkgimport/main.go)

### Lazy Parsing (Memory Optimization)

To reduce memory usage, you can enable lazy parsing which defers the parsing of type declarations until they are actually accessed:

```go
import "github.com/wzshiming/gotype"

// Create an importer with lazy parsing enabled
imp := gotype.NewImporter(gotype.WithLazyParsing())

// Parse source code
scope, err := imp.ImportSource("mypackage", []byte(sourceCode))
if err != nil {
    // handle error
}

// Types are only fully parsed when accessed
myType, ok := scope.ChildByName("MyType")
```

Lazy parsing can reduce memory usage by approximately 25-35% during the initial parsing phase, making it particularly useful when working with large codebases or when only a subset of types needs to be analyzed.

## License

Licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/gotype/blob/master/LICENSE) for the full license text.
