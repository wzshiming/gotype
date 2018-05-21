package gotype

func NewTypeChan(v Type, dir ChanDir) Type {
	return &TypeChan{
		val: v,
		dir: dir,
	}
}

type TypeChan struct {
	typeBase
	dir ChanDir
	val Type
}

func (t *TypeChan) Kind() Kind {
	return Chan
}

func (t *TypeChan) Elem() Type {
	return t.val
}

func (t *TypeChan) ChanDir() ChanDir {
	return t.dir
}
