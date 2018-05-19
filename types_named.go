package gotype

type TypeNamed struct {
	Type
	name string // 名字

}

func (t *TypeNamed) Name() string {
	return t.name
}
