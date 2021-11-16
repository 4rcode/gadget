package gadget

// Inspector returns true only when the provided value can be detected using the
// provided arguments. Otherwise, it returns false and reports the failure.
//
// The arguments can be raw values, or instances of
//   > gadget.Detector
//   > settings.Options
//
// When no arguments are provided, implementations usually expect true as value.
type Inspector func(value interface{}, arguments ...interface{}) bool
