package gotype

import (
	"reflect"
	"sort"
)

type Type interface {
	Name() string
	Kind() Kind
	Key() Type
	Elem() Type // 指针 接口 别名 数组 map值 到下一级
	NumMethods() int
	Methods(int) Type
	NumField() int
	Field(int) Type
	Tag() reflect.StructTag  // field
	Len() int                // 数组
	NumOut() int             // func
	Out(int) Type            // func
	NumIn() int              // func
	In(int) Type             // func
	ChanDir() ChanDir        // chan
	Child(int) Type          // scope
	NumChild() int           // scope
	ChildByName(string) Type // scope
}

type Types []Type

func (t *Types) Add(n Type) {
	if n == nil {
		return
	}
	i := t.SearchIndex(n.Name())
	*t = append(*t, nil)
	l := len(*t)
	copy((*t)[i+1:l], (*t)[i:l-1])
	(*t)[i] = n
}

func (t *Types) Search(name string) Type {
	i := t.SearchIndex(name)
	if i == 0 {
		return nil
	}
	i--
	tt := t.Index(i)
	if tt == nil || tt.Name() != name {
		return nil
	}
	return tt
}

func (t *Types) SearchIndex(name string) int {
	i := sort.Search(t.Len(), func(i int) bool {
		return t.Index(i).Name() < name
	})
	return i
}

func (t *Types) Index(i int) Type {
	if i >= t.Len() {
		return nil
	}
	return (*t)[i]
}

func (t *Types) Len() int {
	return len(*t)
}

type typeBase struct{}

func (t *typeBase) Name() string            { return "" }
func (t *typeBase) Key() Type               { return nil }
func (t *typeBase) Elem() Type              { return nil }
func (t *typeBase) NumField() int           { return 0 }
func (t *typeBase) Field(int) Type          { return nil }
func (t *typeBase) Tag() reflect.StructTag  { return reflect.StructTag("") }
func (t *typeBase) Len() int                { return 0 }
func (t *typeBase) NumOut() int             { return 0 }
func (t *typeBase) Out(int) Type            { return nil }
func (t *typeBase) NumIn() int              { return 0 }
func (t *typeBase) In(int) Type             { return nil }
func (t *typeBase) NumMethods() int         { return 0 }
func (t *typeBase) Methods(int) Type        { return nil }
func (t *typeBase) ChanDir() ChanDir        { return 0 }
func (t *typeBase) Child(int) Type          { return nil }
func (t *typeBase) NumChild() int           { return 0 }
func (t *typeBase) ChildByName(string) Type { return nil }
