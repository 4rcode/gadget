// Code generated by optgen. DO NOT EDIT.
package settings

type LikeFormat string

func (value LikeFormat) ApplyTo(entity interface{}) {
	if target, ok := entity.(interface {
		SetLikeFormat(LikeFormat)
	}); ok {
		target.SetLikeFormat(value)
	}
}

func (target *LikeFormat) SetLikeFormat(value LikeFormat) {
	*target = value
}
