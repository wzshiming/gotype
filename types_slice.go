package gotype

func newTypeSlice(v Type) Type {
	return &typeSlice{
		val: v,
	}
}

type typeSlice struct {
	typeBase
	val Type
}

func (t *typeSlice) Kind() Kind {
	return Slice
}

func (t *typeSlice) Elem() Type {
	return t.val
}
