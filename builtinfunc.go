package gotype

//go:generate stringer -type builtinfunc builtinfunc.go
type builtinfunc uint8

// 内置函数 的参数计算
const (
	_                    builtinfunc = iota
	builtinfuncInt                   // 默认 int
	builtinfuncPtrItem               // 第一个参数的指针
	builtinfuncItem                  // 第一个参数
	builtinfuncInterface             // 接口
	builtinfuncVoid                  // 无
)
