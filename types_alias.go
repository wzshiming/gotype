package gotype

func newTypeAlias(name string, typ Type, info *infoFile) Type {
	return &typeAlias{
		name: name,
		info: info,
		Type: typ,
	}
}

type typeAlias struct {
	name string
	info *infoFile
	Type
}

func (t *typeAlias) Name() string {
	return t.name
}

func (t *typeAlias) String() string {
	return t.name
}

func (t *typeAlias) PkgPath() string {
	return t.info.PkgPath
}
