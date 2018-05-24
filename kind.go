package gotype

//go:generate stringer -type Kind kind.go
type Kind uint8

const (
	Invalid Kind = iota

	// 内置基础类型
	predeclaredTypesBeg
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	String
	Byte
	Rune
	Error
	predeclaredTypesEnd

	// 内置组合类型
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	Struct

	// 特殊
	Field // Struct Field
	Var   // 变量名
	Scope // 包名 和 函数内作用域
)
