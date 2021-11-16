// Code generated by optgen. DO NOT EDIT.
package settings

type Fallback []interface{}

func (value Fallback) ApplyTo(entity interface{}) {
	if target, ok := entity.(interface {
		SetFallback(Fallback)
	}); ok {
		target.SetFallback(value)
	}
}

func (target *Fallback) SetFallback(value Fallback) {
	if len(value) < 1 {
		*target = nil
	} else {
		*target = append(*target, value...)
	}
}
