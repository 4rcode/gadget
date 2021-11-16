package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/settings"
)

type MultiDetector struct {
	Breaker, Inverted bool

	detectors []gadget.Detector
}

func (d MultiDetector) Build(values ...interface{}) gadget.Detector {
	return d.Values(values[0], values[1:]...)
}

func (d MultiDetector) Detect(value interface{}, options gadget.Options) error {
	cfg := struct {
		settings.NegativeFormat
		settings.Separator
	}{
		settings.NegativeFormat("not %[1]v"),
		settings.Separator(" | "),
	}

	options.ApplyToAll(&cfg)

	for _, detector := range d.detectors {
		err := detector.Detect(value, options)

		if err == nil != d.Breaker {
			continue
		}

		if d.Breaker != d.Inverted {
			return nil
		}

		if !d.Inverted {
			return err
		}

		err = detector.Detect(gadget.Fake{}, options)

		return d.not(cfg.NegativeFormat, err)
	}

	if d.Breaker == d.Inverted {
		return nil
	}

	items := make([]string, 0, len(d.detectors))

	for _, detector := range d.detectors {
		err := detector.Detect(gadget.Fake{}, options)

		if d.Inverted {
			err = d.not(cfg.NegativeFormat, err)
		}

		items = append(items, err.Error())
	}

	text := strings.Join(items, string(cfg.Separator))

	return errors.New(text)
}

func (d MultiDetector) not(format settings.NegativeFormat, err error) error {
	return fmt.Errorf(string(format), err, "")
}

func (d MultiDetector) Values(value interface{}, values ...interface{}) MultiDetector {
	values = append([]interface{}{value}, values...)

	for _, value := range values {
		detector, ok := value.(gadget.Detector)

		if !ok {
			detector = equalTo(value)
		}

		d.detectors = append(d.detectors, detector)
	}

	return d
}
