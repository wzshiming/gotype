package gotype

// info holds result type information.
// Only the information for which a map is provided is collected.
// If the package has type errors, the collected information may be incomplete.
type info struct {
	Named   types            // var, func, type, packgae
	Methods map[string]types // type method
	Src     string
	PkgPath string
	Goroot  bool
}

func newInfo(src string, pkg string, goroot bool) *info {
	return &info{
		Src:     src,
		PkgPath: pkg,
		Goroot:  goroot,
		Methods: map[string]types{},
	}
}
