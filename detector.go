package gadget

// Detector is responsible for detecting a value.
//
// Implementations must return an error when a Fake value is provided.
type Detector interface {
	// Detect returns an error when the provided value can NOT be detected, nil
	// otherwise. Options are provided to modify or enhance the detection error.
	Detect(value interface{}, options Options) error
}

// DetectorFunc is responsible for converting a function into a Detector.
type DetectorFunc func(interface{}, Options) error

func (f DetectorFunc) Detect(value interface{}, options Options) error {
	return f(value, options)
}

// Fake can NOT be detected and, when provided as value, it must make any
// detector fail.
type Fake struct{}

func (Fake) String() string {
	return "Fake"
}
