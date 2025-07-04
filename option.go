package assert

import (
	"fmt"
	"testing"
)

type options struct {
	messageTransforms []func(msg string) string
	report            ReportFunc
}

func buildOptions(opts []Option) *options {
	o := &options{
		report: testing.TB.Fatal,
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Option is an option for an assertion.
type Option func(*options)

// Lazy returns an [Option] that defers the evaluation of the option.
//
// It helps to reduce allocations when the option is not used.
func Lazy(f func() Option) Option {
	return func(o *options) {
		f()(o)
	}
}

// Options returns an [Option] that combines several options.
func Options(opts ...Option) Option {
	return func(o *options) {
		for _, opt := range opts {
			opt(o)
		}
	}
}

// MessageTransform returns an [Option] that adds a message transform function.
// The function is called before the [ReportFunc].
// If several function are added, they're called in order.
func MessageTransform(f func(msg string) string) Option {
	return func(o *options) {
		o.messageTransforms = append(o.messageTransforms, f)
	}
}

// Message returns an [Option] that sets the message.
func Message(msg string) Option {
	return MessageTransform(func(_ string) string {
		return msg
	})
}

// Messagef returns an [Option] that sets the formatted message.
func Messagef(format string, args ...any) Option {
	return MessageTransform(func(_ string) string {
		return fmt.Sprintf(format, args...)
	})
}

// MessageWrap returns an [Option] that wraps the message.
// The final message is "<msg>: <original message>".
func MessageWrap(msg string) Option {
	return MessageTransform(func(wrappedMsg string) string {
		return msg + ": " + wrappedMsg
	})
}

// MessageWrapf returns an [Option] that wraps the message.
// The final message is "<format msg>: <original message>".
func MessageWrapf(format string, args ...any) Option {
	return MessageTransform(func(wrappedMsg string) string {
		return fmt.Sprintf(format, args...) + ": " + wrappedMsg
	})
}

// ReportFunc is a function that is called when an assertion fails.
//
// It is implemented by [testing.TB.Fatal]|[testing.TB.Error]|[testing.TB.Skip]|[testing.TB.Log].
// The default value is [testing.TB.Fatal].
type ReportFunc func(tb testing.TB, args ...any)

// Report returns an [Option] that sets the [ReportFunc].
func Report(f ReportFunc) Option {
	return func(o *options) {
		o.report = f
	}
}

// ReportFatal returns an [Option] that sets the [ReportFunc] to [testing.TB.Fatal].
func ReportFatal() Option {
	return Report(testing.TB.Fatal)
}

// ReportError returns an [Option] that sets the [ReportFunc] to [testing.TB.Error].
func ReportError() Option {
	return Report(testing.TB.Error)
}

// ReportSkip returns an [Option] that sets the [ReportFunc] to [testing.TB.Skip].
func ReportSkip() Option {
	return Report(testing.TB.Skip)
}

// ReportLog returns an [Option] that sets the [ReportFunc] to [testing.TB.Log].
func ReportLog() Option {
	return Report(testing.TB.Log)
}
