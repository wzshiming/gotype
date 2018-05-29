package gotype

import "fmt"

func newTypeMap(k, v Type) Type {
	return &typeMap{
		key:  k,
		elem: v,
	}
}

type typeMap struct {
	typeBase
	key, elem Type
}

func (t *typeMap) String() string {
	return fmt.Sprintf("map[%v]%v", t.key, t.elem)
}

func (t *typeMap) Kind() Kind {
	return Map
}

func (t *typeMap) Key() Type {
	return t.key
}

func (t *typeMap) Elem() Type {
	return t.elem
}
