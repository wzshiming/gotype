package gotype

import "fmt"

func newTypeMap(k, v Type) Type {
	return &typeMap{
		key: k,
		val: v,
	}
}

type typeMap struct {
	typeBase
	key, val Type
}

func (t *typeMap) String() string {
	return fmt.Sprintf("map[%v]%v", t.key, t.val)
}

func (t *typeMap) Kind() Kind {
	return Map
}

func (t *typeMap) Key() Type {
	return t.key
}

func (t *typeMap) Elem() Type {
	return t.val
}
