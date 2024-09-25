package gotype

//go:generate stringer -type Kind kind.go

// Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type Kind uint8

// Define kind
const (
	Invalid Kind = iota

	// Built-in base type

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
	Any
	Comparable
	predeclaredTypesEnd

	// Built-in combination

	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	Struct

	// Special is different from other Kinds

	Field       // a Struct Field
	Scope       // package or func body
	Declaration // a top-level function, variable, or constant.
	Param       // a Type Param
)
