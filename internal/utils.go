package internal

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/settings"
)

func Comparable(value interface{}) bool {
	nature := reflect.TypeOf(value)

	if nature == nil {
		return true
	}

	return nature.Comparable()
}

func Convert(value, target interface{}) interface{} {
	nature := reflect.TypeOf(target)

	if nature == nil {
		return value
	}

	reference := reflect.ValueOf(value)

	if !reference.IsValid() {
		return value
	}

	if !reference.CanConvert(nature) {
		return value
	}

	reference = reference.Convert(nature)

	return reference.Interface()
}

func IsFake(value interface{}) bool {
	return value == gadget.Fake{}
}

func Print(options gadget.Options, value interface{}) string {
	reference := reflect.ValueOf(value)

	switch reference.Kind() {
	case reflect.Func:
		return runtime.FuncForPC(
			reference.Pointer(),
		).Name()
	}

	format := settings.Format("%#[1]v")

	options.ApplyToAll(&format)

	return fmt.Sprintf(string(format), value, "")
}
