# gotype

Golang 源代码解析，像反射包一样使用

[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/gotype)](https://goreportcard.com/report/github.com/wzshiming/gotype)
[![GoDoc](https://godoc.org/github.com/wzshiming/gotype?status.svg)](https://godoc.org/github.com/wzshiming/gotype)
[![GitHub license](https://img.shields.io/github/license/wzshiming/gotype.svg)](https://github.com/wzshiming/gotype/blob/master/LICENSE)

- [English](https://github.com/wzshiming/gotype/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/gotype/blob/master/README_cn.md)

## 用法

[API 文档](https://godoc.org/github.com/wzshiming/gotype)

[示例](https://github.com/wzshiming/gotype/blob/master/cmd/pkgimport/main.go)

### 惰性解析（内存优化）

为了减少内存使用，您可以启用惰性解析，它会推迟类型声明的解析，直到实际访问时才进行：

```go
import "github.com/wzshiming/gotype"

// 创建一个启用惰性解析的导入器
imp := gotype.NewImporter(gotype.WithLazyParsing())

// 解析源代码
scope, err := imp.ImportSource("mypackage", []byte(sourceCode))
if err != nil {
    // 处理错误
}

// 类型只在访问时才被完全解析
myType, ok := scope.ChildByName("MyType")
```

惰性解析可以在初始解析阶段减少约 25-35% 的内存使用，这在处理大型代码库或只需要分析类型子集时特别有用。

## 许可证

软包根据MIT License。有关完整的许可证文本，请参阅[LICENSE](https://github.com/wzshiming/gotype/blob/master/LICENSE)。
