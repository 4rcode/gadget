package gadget

// Builder is responsible for creating the main Detector used by an Inspector.
type Builder interface {
	// Build returns a new Detector given a sequence of arguments.
	Build(arguments ...interface{}) Detector
}

// BuilderFunc is responsible for converting a function into a Builder.
type BuilderFunc func(...interface{}) Detector

func (f BuilderFunc) Build(arguments ...interface{}) Detector {
	return f(arguments...)
}
