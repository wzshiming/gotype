package gotype

//go:generate stringer -type builtinfunc builtinfunc.go

type builtinfunc uint8

// Built-in function results
const (
	_                    builtinfunc = iota
	builtinfuncInt                   // Returns int
	builtinfuncPtrItem               // Returns the type pointer of the first parameter
	builtinfuncItem                  // Returns the first parameter
	builtinfuncInterface             // Returns interface
	builtinfuncVoid                  // Returns void
)
