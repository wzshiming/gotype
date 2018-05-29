package gotype

import (
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

	// Name returns the type's name within its package.
	// It returns an empty string for unnamed types.
	Name() string

	// Kind returns the specific kind of this type.
	Kind() Kind

	// Key returns a map type's key type.
	// It panics if the type's Kind is not Map.
	Key() Type

	// Elem returns a type's element type.
	// It panics if the type's Kind is not Var, Array, Chan, Map, Ptr, or Slice.
	Elem() Type

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
	// Not contain anonymo
	// It panics if the type's Kind is not Struct.
	// It panics if i is not in the range [0, NumField()).
	Field(int) Type

	// FieldByName returns the struct field with the given name
	// and a boolean indicating if the field was found.
	FieldByName(string) (Type, bool)

	// NumField returns a struct type's field count.
	// Not contain anonymo
	// It panics if the type's Kind is not Struct.
	NumField() int

	// Method returns the i'th method in the type's method set.
	// Not contain anonymo
	// It panics if i is not in the range [0, NumMethod()).
	//
	// For a non-interface type T or *T, the returned Method's Type and Func
	// fields describe a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	Methods(int) Type

	// MethodByName returns the method with that name in the type's
	// method set and a boolean indicating if the method was found.
	//
	// For a non-interface type T or *T, the returned Method's Type and Func
	// fields describe a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	MethodsByName(string) (Type, bool)

	// NumMethod returns the number of exported methods in the type's method set.
	// Not contain anonymo
	NumMethods() int

	// Field returns a scope type's i'th field.
	// It panics if i is not in the range [0, NumChild()).
	Child(int) Type

	// ChildByName returns the scope with the given name
	// and a boolean indicating if the field was found.
	ChildByName(string) (Type, bool)

	// NumChild returns a scope type's field count.
	NumChild() int

	// Anonymo returns the i'th type in the type's anonymo set.
	// It panics if i is not in the range [0, NumAnonymo()).
	// It panics if the type's Kind is not Interface or Struct.
	Anonymo(int) Type

	// AnonymoByName returns the anonymo type with the given name
	// and a boolean indicating if the field was found.
	AnonymoByName(string) (Type, bool)

	// NumChild returns a anonymo type's field count.
	// It panics if the type's Kind is not Interface or Struct.
	NumAnonymo() int
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
	tt := t.Index(i - 1)
	if tt == nil || tt.Name() != name {
		t.add(i, n)
	}
	return
}

func (t *types) Search(name string) (Type, bool) {
	i := t.SearchIndex(name)
	if i == 0 {
		return nil, false
	}
	tt := t.Index(i - 1)
	if tt == nil || tt.Name() != name {
		return nil, false
	}
	return tt, true
}

func (t *types) SearchIndex(name string) int {
	i := sort.Search(t.Len(), func(i int) bool {
		d := t.Index(i)
		if d == nil {
			return false
		}
		return d.Name() < name
	})
	return i
}

func (t *types) Index(i int) Type {
	if i >= t.Len() || i < 0 {
		return nil
	}
	return (*t)[i]
}

func (t *types) Len() int {
	return len(*t)
}

type typeBase struct{}

func (t *typeBase) Name() string {
	return ""
}

func (t *typeBase) Kind() Kind {
	return Invalid
}

func (t *typeBase) Key() Type {
	panic("Key of non-map type")
	return nil
}

func (t *typeBase) Elem() Type {
	panic("Elem of invalid type")
	return nil
}

func (t *typeBase) Tag() reflect.StructTag {
	panic("Tag of non-field type")
	return ""
}

func (t *typeBase) Len() int {
	panic("Len of non-array type")
	return 0
}

func (t *typeBase) Value() string {
	return ""
}

func (t *typeBase) ChanDir() ChanDir {
	panic("ChanDir of non-chan type")
	return 0
}

func (t *typeBase) Out(int) Type {
	panic("Out of non-func type")
	return nil
}

func (t *typeBase) NumOut() int {
	panic("NumOut of non-func type")
	return 0
}

func (t *typeBase) In(int) Type {
	panic("In of non-func type")
	return nil
}

func (t *typeBase) NumIn() int {
	panic("NumIn of non-func type")
	return 0
}

func (t *typeBase) IsVariadic() bool {
	panic("IsVariadic of non-func type")
	return false
}
func (t *typeBase) Field(int) Type {
	panic("Field of non-struct type")
	return nil
}
func (t *typeBase) FieldByName(string) (Type, bool) {
	panic("FieldByName of non-struct type")
	return nil, false
}
func (t *typeBase) NumField() int {
	panic("NumField of non-struct type")
	return 0
}
func (t *typeBase) Methods(int) Type {
	//panic("Methods of invalid type")
	return nil
}

func (t *typeBase) MethodsByName(string) (Type, bool) {
	//panic("MethodsByName of invalid type")
	return nil, false
}

func (t *typeBase) NumMethods() int {
	//panic("NumMethods of invalid type")
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

func (t *typeBase) Anonymo(int) Type {
	panic("Anonymo of invalid type")
	return nil
}

func (t *typeBase) AnonymoByName(string) (Type, bool) {
	panic("AnonymoByName of invalid type")
	return nil, false
}

func (t *typeBase) NumAnonymo() int {
	panic("NumAnonymo of invalid type")
	return 0
}
