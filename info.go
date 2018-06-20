package gotype

// info holds result type information.
// Only the information for which a map is provided is collected.
// If the package has type errors, the collected information may be incomplete.
type info struct {
	Named   types            // var, func, type, packgae
	Methods map[string]types // type method
	PkgPath string
	Goroot  bool
}

func newInfo(pkg string, goroot bool) *info {
	return &info{
		PkgPath: pkg,
		Goroot:  goroot,
		Methods: map[string]types{},
	}
}
