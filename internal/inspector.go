package internal

import (
	"github.com/4rcode/gadget/settings"
	"github.com/4rcode/gnomic"
)

type Inspector struct {
	gnomic.Options
}

func (i Inspector) Inspect(value interface{}, arguments ...interface{}) bool {
	var factory _factory

	opts, arguments := i.parse(arguments...)
	opts = gnomic.Options{
		settings.Builder{},
	}.Append(i).Append(opts)

	opts.ApplyTo(&factory)

	if len(arguments) < 1 {
		arguments = append(arguments, factory.Fallback...)
	}

	detector := factory.newDetector(arguments...)
	err := detector.Detect(value, opts)

	if err == nil {
		return true
	}

	for _, reporter := range factory.Reporters {
		if reporter == nil {
			continue
		}

		if helper, ok := reporter.(interface {
			Helper()
		}); ok {
			helper.Helper()
		}

		reporter.Report(value, err, opts)
	}

	return false
}

func (Inspector) parse(arguments ...interface{}) (gnomic.Options, []interface{}) {
	var values []interface{}
	var options gnomic.Options

	for _, argument := range arguments {
		if opts, ok := argument.(settings.Options); ok {
			options = options.Append(opts...)
		} else {
			values = append(values, argument)
		}
	}

	return options, values
}
