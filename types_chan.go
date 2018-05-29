package gotype

import "fmt"

func newTypeChan(v Type, dir ChanDir) Type {
	return &typeChan{
		elem: v,
		dir:  dir,
	}
}

type typeChan struct {
	typeBase
	dir  ChanDir
	elem Type
}

func (t *typeChan) String() string {
	switch t.dir {
	case RecvDir:
		return fmt.Sprintf("<-chan %v", t.elem)
	case SendDir:
		return fmt.Sprintf("chan<- %v", t.elem)
	case BothDir:
	}
	return fmt.Sprintf("chan %v", t.elem)
}

func (t *typeChan) Kind() Kind {
	return Chan
}

func (t *typeChan) Elem() Type {
	return t.elem
}

func (t *typeChan) ChanDir() ChanDir {
	return t.dir
}
