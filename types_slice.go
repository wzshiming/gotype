package gotype

func NewTypeSlice(v Type) Type {
	return &TypeSlice{
		val: v,
	}
}

type TypeSlice struct {
	typeBase
	val Type
}

func (t *TypeSlice) Kind() Kind {
	return Slice
}

func (t *TypeSlice) Elem() Type {
	return t.val
}
