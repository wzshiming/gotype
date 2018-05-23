package gotype

import "fmt"

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

func (t *typeChan) String() string {
	switch t.dir {
	case RecvDir:
		return fmt.Sprintf("<-chan %v", t.val)
	case SendDir:
		return fmt.Sprintf("chan<- %v", t.val)
	case BothDir:
	}
	return fmt.Sprintf("chan %v", t.val)
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
