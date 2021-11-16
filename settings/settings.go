// - Builder configures the Builder used by an Inspector.
//
//
// - Fallback configures arguments to use, when none are supplied during an
// assertion. In this way, you can replace the default Detector with one or more
// of your choice.
//
//
// - Format configures how to print Detector errors (e.g. EqualTo()).
//
// e.g.
//   "[ %#[1]v ]"
// where
//   %[1]v = value
//   %[2]v = ""
//
// The second argument is an empty string you can use to hide the value. It can
// also be used as a spacer.
//
//
// - LikeFormat configures how to print the Like() Detector errors.
//
//
// - NegativeFormat configures how to print the Not() Detector errors.
//
//
// - Reporters configures the Reporter instances used by an Inspector.
//
//
// - Separator configures the string used for printing multiple delimited
// Detector instances.
//
//
// - StringFormat configures how to print the String() Detector errors.
//
//
// - Template configures the string used by Reporter instances to print their
// message.
//
// e.g.
//   "%[1]v != %[2]v"
// where
//   %[1]v = value
//   %[2]v = target
//   %[3]v = ""
//
// The third argument is an empty string you can use to hide both value and
// target. It can also be used as a spacer.
package settings

//go:generate optgen -n Builder -t [1]gadget.Builder -i="github.com/4rcode/gadget"
//go:generate optgen -n Fallback -t []interface{} -a -c nil
//go:generate optgen -n Reporters -t []gadget.Reporter -a -c nil -i="github.com/4rcode/gadget"
//go:generate optgen -t string -n Format
//go:generate optgen -t string -n LikeFormat
//go:generate optgen -t string -n NegativeFormat
//go:generate optgen -t string -n Separator
//go:generate optgen -t string -n StringFormat
//go:generate optgen -t string -n Template

import (
	"github.com/4rcode/gnomic"
)

// Options is a container for any option.
type Options []gnomic.Option
