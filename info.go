package gotype

// info holds result type information.
// Only the information for which a map is provided is collected.
// If the package has type errors, the collected information may be incomplete.
type info struct {
	PkgNamed map[string]types // map[file]packgae
	Named    types            // var, func, type,
	Methods  map[string]types // map[type]method
	Src      string
	PkgPath  string
	Goroot   bool
}

type infoFile struct {
	*info
	filename string
}

func newInfo(src string, pkg string, goroot bool) *info {
	return &info{
		Src:      src,
		PkgPath:  pkg,
		Goroot:   goroot,
		Methods:  map[string]types{},
		PkgNamed: map[string]types{},
	}
}

func (i *info) File(filename string) *infoFile {
	return &infoFile{
		info:     i,
		filename: filename,
	}
}

func (i *infoFile) GetPkgOrType(name string) (Type, bool) {
	if pkg, ok := i.PkgNamed[i.filename]; ok {
		t, ok := pkg.Search(name)
		if ok {
			return t, ok
		}
	}
	if i.filename == "" {
		for _, pkg := range i.PkgNamed {
			t, ok := pkg.Search(name)
			if ok {
				return t, ok
			}
		}
	}
	return i.Named.Search(name)
}

func (i *infoFile) AddPkg(t Type) {
	pkg, _ := i.PkgNamed[i.filename]
	pkg.Add(t)
	i.PkgNamed[i.filename] = pkg
}

func (i *infoFile) AddType(t Type) {
	i.Named.Add(t)
}

func (i *infoFile) AddMethod(name string, t Type) {
	methods := i.Methods[name]
	methods.Add(t)
	i.Methods[name] = methods
}
