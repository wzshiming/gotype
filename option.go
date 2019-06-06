package gotype

// Option some basic options
type Option func(i *Importer)

// ErrorHandler returns the error handler option
func ErrorHandler(f func(error)) Option {
	return func(i *Importer) {
		i.errorHandler = f
	}
}

// WithCommentLocator sets comment locator
func WithCommentLocator() Option {
	return func(i *Importer) {
		i.isCommentLocator = true
	}
}

// ImportHandler returns the import handler option
func ImportHandler(f func(path, src, dir string)) Option {
	return func(i *Importer) {
		i.importHandler = f
	}
}
