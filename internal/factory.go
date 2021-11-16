package internal

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/settings"
)

type _factory struct {
	gadget.Builder
	settings.Fallback
	settings.Reporters
}

func (f *_factory) newDetector(arguments ...interface{}) gadget.Detector {
	if len(arguments) < 1 {
		arguments = append(arguments, equalTo(true))
	}

	detector := f.Builder.Build(arguments...)

	if detector == nil {
		detector = equalTo(true)
	}

	return detector
}

func (f *_factory) SetBuilder(builder settings.Builder) {
	if builder[0] == nil {
		f.Builder = MultiDetector{}
	} else {
		f.Builder = builder[0]
	}
}
