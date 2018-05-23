package gotype

import (
	"reflect"
	"sort"
)

type Type interface {
	Name() string
	Kind() Kind
	Key() Type                 // map key
	Elem() Type                // map value, ptr, slice, array, chan
	Tag() reflect.StructTag    // field
	Len() int                  // array
	ChanDir() ChanDir          // chan
	Out(int) Type              // func
	NumOut() int               // func
	In(int) Type               // func
	NumIn() int                // func
	IsVariadic() bool          // func
	Field(int) Type            // struct
	FieldByName(string) Type   // struct
	NumField() int             // struct
	Methods(int) Type          // named, alias, interface
	MethodsByName(string) Type // named, alias, interface
	NumMethods() int           // named, alias, interface
	Child(int) Type            // scope
	ChildByName(string) Type   // scope
	NumChild() int             // scope
	Anonymo(int) Type          // struct, interface
	AnonymoByName(string) Type // struct, interface
	NumAnonymo() int           // struct, interface
}

type Types []Type

func (t *Types) add(i int, n Type) {
	*t = append(*t, n)
	l := len(*t)
	copy((*t)[i+1:l], (*t)[i:l-1])
	(*t)[i] = n
}

func (t *Types) Add(n Type) {
	if n == nil {
		return
	}
	name := n.Name()
	i := t.SearchIndex(name)
	t.add(i, n)
}

func (t *Types) AddNoRepeat(n Type) {
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

func (t *Types) Search(name string) Type {
	i := t.SearchIndex(name)
	if i == 0 {
		return nil
	}
	tt := t.Index(i - 1)
	if tt == nil || tt.Name() != name {
		return nil
	}
	return tt
}

func (t *Types) SearchIndex(name string) int {
	i := sort.Search(t.Len(), func(i int) bool {
		d := t.Index(i)
		if d == nil {
			return false
		}
		return d.Name() < name
	})
	return i
}

func (t *Types) Index(i int) Type {
	if i >= t.Len() || i < 0 {
		return nil
	}
	return (*t)[i]
}

func (t *Types) Len() int {
	return len(*t)
}

type typeBase struct{}

func (t *typeBase) Name() string              { return "" }
func (t *typeBase) Kind() Kind                { return Invalid }
func (t *typeBase) Key() Type                 { return nil }
func (t *typeBase) Elem() Type                { return nil }
func (t *typeBase) Tag() reflect.StructTag    { return reflect.StructTag("") }
func (t *typeBase) Len() int                  { return 0 }
func (t *typeBase) ChanDir() ChanDir          { return 0 }
func (t *typeBase) Out(int) Type              { return nil }
func (t *typeBase) NumOut() int               { return 0 }
func (t *typeBase) In(int) Type               { return nil }
func (t *typeBase) NumIn() int                { return 0 }
func (t *typeBase) IsVariadic() bool          { return false }
func (t *typeBase) Field(int) Type            { return nil }
func (t *typeBase) FieldByName(string) Type   { return nil }
func (t *typeBase) NumField() int             { return 0 }
func (t *typeBase) Methods(int) Type          { return nil }
func (t *typeBase) MethodsByName(string) Type { return nil }
func (t *typeBase) NumMethods() int           { return 0 }
func (t *typeBase) Child(int) Type            { return nil }
func (t *typeBase) ChildByName(string) Type   { return nil }
func (t *typeBase) NumChild() int             { return 0 }
func (t *typeBase) Anonymo(int) Type          { return nil }
func (t *typeBase) AnonymoByName(string) Type { return nil }
func (t *typeBase) NumAnonymo() int           { return 0 }
