package gotype

func NewTypeMap(k, v Type) Type {
	return &TypeMap{
		key: k,
		val: v,
	}
}

type TypeMap struct {
	typeBase
	key, val Type
}

func (t *TypeMap) Kind() Kind {
	return Map
}

func (t *TypeMap) Key() Type {
	return t.key
}

func (t *TypeMap) Elem() Type {
	return t.val
}
