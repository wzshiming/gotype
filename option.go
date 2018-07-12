package gotype

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
