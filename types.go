package gotype

import (
	"go/ast"
	"reflect"
	"sort"
)

// Type is the representation of a Go type.
//
// Not all methods apply to all kinds of types. Restrictions,
// if any, are noted in the documentation for each method.
// Use the Kind method to find out the kind of type before
// calling kind-specific methods. Calling a method
// inappropriate to the kind of type causes a run-time panic.
//
// Type values are comparable, such as with the == operator,
// so they can be used as map keys.
// Two Type values are equal if they represent identical types.
type Type interface {

	// String returns a string representation of the type.
	// The string representation may use shortened package names
	// (e.g., base64 instead of "encoding/base64") and is not
	// guaranteed to be unique among types. To test for type identity,
	// compare the Types directly.
	String() string

	// PkgPath returns a named type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
	// the package path will be the empty string.
	PkgPath() string

	// IsGoroot returns package is found in Go root
	IsGoroot() bool

	// Name returns the type's name within its package.
	// It returns an empty string for unnamed types.
	Name() string

	// Kind returns the specific kind of this type.
	Kind() Kind

	// Key returns a map type's key type.
	// It panics if the type's Kind is not Map.
	Key() Type

	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	Elem() Type

	// Declaration returns a type's declaration.
	// It panics if the type's Kind is not declaration.
	Declaration() Type

	// Tag returns a field type's tag.
	// It panics if the type's Kind is not Field.
	Tag() reflect.StructTag

	// Len returns an array type's length.
	// It panics if the type's Kind is not Array.
	Len() int

	// Value returns a type's value.
	// Only get constants
	Value() string

	// ChanDir returns a channel type's direction.
	// It panics if the type's Kind is not Chan.
	ChanDir() ChanDir

	// Out returns the type of a function type's i'th output parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumOut()).
	Out(int) Type

	// NumOut returns a function type's output parameter count.
	// It panics if the type's Kind is not Func.
	NumOut() int

	// In returns the type of a function type's i'th input parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumIn()).
	In(int) Type

	// NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	NumIn() int

	// IsVariadic reports whether a function type's final input parameter
	// is a "..." parameter. If so, t.In(t.NumIn() - 1) returns the parameter's
	// implicit actual type []T.
	//
	// For concreteness, if t represents func(x int, y ... float64), then
	//
	//	t.NumIn() == 2
	//	t.In(0) is the reflect.Type for "int"
	//	t.In(1) is the reflect.Type for "[]float64"
	//	t.IsVariadic() == true
	//
	// IsVariadic panics if the type's Kind is not Func.
	IsVariadic() bool

	// Field returns a struct type's i'th field.
	// It panics if the type's Kind is not Struct.
	// It panics if i is not in the range [0, NumField()).
	Field(int) Type

	// FieldByName returns the struct field with the given name
	// and a boolean indicating if the field was found.
	FieldByName(string) (Type, bool)

	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	NumField() int

	// IsAnonymous returns is an embedded field
	// It panics if the type's Kind is not Field.
	IsAnonymous() bool

	// Method returns the i'th method in the type's method set.
	// Not contain anonymo
	// It panics if i is not in the range [0, NumMethod()).
	//
	// For a non-interface type T or *T, the returned Method's Type and Func
	// fields describe a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	Method(int) Type

	// MethodByName returns the method with that name in the type's
	// method set and a boolean indicating if the method was found.
	//
	// For a non-interface type T or *T, the returned Method's Type and Func
	// fields describe a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	MethodByName(string) (Type, bool)

	// NumMethod returns the number of exported methods in the type's method set.
	// Not contain anonymo
	NumMethod() int

	// Child returns a scope type's i'th child.
	// It panics if i is not in the range [0, NumChild()).
	Child(int) Type

	// ChildByName returns the scope with the given name
	// and a boolean indicating if the child was found.
	ChildByName(string) (Type, bool)

	// NumChild returns a scope type's child count.
	NumChild() int

	// Origin returns the type's origin data within its package.
	Origin() ast.Node

	// Doc returns the type's doc within its package.
	Doc() *ast.CommentGroup

	// Comment returns the type's comment within its package.
	Comment() *ast.CommentGroup

	// NumParam returns a type's param count.
	NumParam() int

	// Param returns a type's i'th param.
	Param(int) Type

	// ParamByName returns the type's param with the given name
	// and a boolean indicating if the field was found.
	ParamByName(string) (Type, bool)
}

type types []Type

func (t *types) add(i int, n Type) {
	*t = append(*t, n)
	l := len(*t)
	copy((*t)[i+1:l], (*t)[i:l-1])
	(*t)[i] = n
}

func (t *types) Add(n Type) {
	if n == nil {
		return
	}
	name := n.Name()
	i := t.SearchIndex(name)
	t.add(i, n)
}

func (t *types) AddNoRepeat(n Type) {
	if n == nil {
		return
	}
	name := n.Name()
	i := t.SearchIndex(name)
	tt := t.Index(i)
	if tt == nil || tt.Name() != name {
		t.add(i, n)
	}
	return
}

func (t *types) Search(name string) (Type, bool) {
	i := t.SearchIndex(name)
	tt := t.Index(i)
	if tt == nil || tt.Name() != name {
		return nil, false
	}
	return tt, true
}

func (t *types) SearchIndex(name string) int {
	return sort.Search(t.Len(), func(i int) bool {
		return t.Index(i).Name() <= name
	})
}

func (t *types) Index(i int) Type {
	if i >= t.Len() {
		return nil
	}
	return (*t)[i]
}

func (t *types) Len() int {
	return len(*t)
}

type typeBase struct{}

func (t *typeBase) IsGoroot() bool {
	return false
}

func (t *typeBase) PkgPath() string {
	return ""
}

func (t *typeBase) Name() string {
	return ""
}

func (t *typeBase) Kind() Kind {
	return Invalid
}

func (t *typeBase) Key() Type {
	panic("Key of non-map type")
}

func (t *typeBase) Elem() Type {
	panic("Elem of invalid type")
}

func (t *typeBase) Declaration() Type {
	panic("Declaration of non-declaration type")
}

func (t *typeBase) Tag() reflect.StructTag {
	panic("Tag of non-field type")
}

func (t *typeBase) Len() int {
	panic("Len of non-array type")
}

func (t *typeBase) Value() string {
	return ""
}

func (t *typeBase) ChanDir() ChanDir {
	panic("ChanDir of non-chan type")
}

func (t *typeBase) Out(int) Type {
	panic("Out of non-func type")
}

func (t *typeBase) NumOut() int {
	panic("NumOut of non-func type")
}

func (t *typeBase) In(int) Type {
	panic("In of non-func type")
}

func (t *typeBase) NumIn() int {
	panic("NumIn of non-func type")
}

func (t *typeBase) IsVariadic() bool {
	panic("IsVariadic of non-func type")
}

func (t *typeBase) Field(int) Type {
	panic("Field of non-struct type")
}

func (t *typeBase) FieldByName(string) (Type, bool) {
	panic("FieldByName of non-struct type")
}

func (t *typeBase) NumField() int {
	panic("NumField of non-struct type")
}

func (t *typeBase) IsAnonymous() bool {
	panic("IsAnonymous of non-field type")
}

func (t *typeBase) Method(int) Type {
	//panic("Method of invalid type")
	return nil
}

func (t *typeBase) MethodByName(string) (Type, bool) {
	//panic("MethodByName of invalid type")
	return nil, false
}

func (t *typeBase) NumMethod() int {
	//panic("NumMethod of invalid type")
	return 0
}

func (t *typeBase) Child(int) Type {
	//panic("Child of invalid type")
	return nil
}

func (t *typeBase) ChildByName(string) (Type, bool) {
	//panic("ChildByName of invalid type")
	return nil, false
}

func (t *typeBase) NumChild() int {
	//panic("NumChild of invalid type")
	return 0
}

func (t *typeBase) Origin() ast.Node {
	return nil
}

func (t *typeBase) Doc() *ast.CommentGroup {
	return nil
}

func (t *typeBase) Comment() *ast.CommentGroup {
	return nil
}

func (t *typeBase) NumParam() int {
	return 0
}

func (t *typeBase) Param(int) Type {
	panic("Param of non-param type")
}

func (t *typeBase) ParamByName(string) (Type, bool) {
	panic("ParamByName of non-param type")
}
