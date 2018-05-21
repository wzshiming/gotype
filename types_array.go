package gotype

func NewTypeArray(v Type, l int) Type {
	return &TypeArray{
		val: v,
		le:  l,
	}
}

type TypeArray struct {
	typeBase
	le  int
	val Type
}

func (t *TypeArray) Kind() Kind {
	return Array
}

func (t *TypeArray) Elem() Type {
	return t.val
}

func (t *TypeArray) Len() int {
	return t.le
}
