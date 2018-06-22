package gotype

type option func(i *Importer)

// ErrorHandler returns the error handler option
func ErrorHandler(f func(error)) option {
	return func(i *Importer) {
		i.errorHandler = f
	}
}

// WithCommentLocator sets comment locator
func WithCommentLocator() option {
	return func(i *Importer) {
		i.isCommentLocator = true
	}
}
