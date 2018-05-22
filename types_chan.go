package gotype

func newTypeChan(v Type, dir ChanDir) Type {
	return &typeChan{
		val: v,
		dir: dir,
	}
}

type typeChan struct {
	typeBase
	dir ChanDir
	val Type
}

func (t *typeChan) Kind() Kind {
	return Chan
}

func (t *typeChan) Elem() Type {
	return t.val
}

func (t *typeChan) ChanDir() ChanDir {
	return t.dir
}
