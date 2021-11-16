package gadget

// Reporter is responsible for reporting a Detector error.
type Reporter interface {
	// Report implementations usually print a message, accepting as arguments the
	// tested value, the detection error and the configured Options.
	Report(value interface{}, error error, options Options)
}

// ReporterFunc is responsible for converting a function into a Reporter.
type ReporterFunc func(interface{}, error, Options)

func (f ReporterFunc) Report(value interface{}, err error, options Options) {
	f(value, err, options)
}
