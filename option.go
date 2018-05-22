package gotype

type option func(i *Importer)

func ErrorHandler(f func(error)) option {
	return func(i *Importer) {
		i.errorHandler = f
	}
}
