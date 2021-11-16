package gadget

// Options is a collection of configured values.
type Options interface {
	// ApplyTo is responsible for applying each configured value to all targets.
	ApplyToAll(targets ...interface{})
}
