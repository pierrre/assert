package assert

import (
	"fmt"
	"testing"
)

type options struct {
	messageTransforms []func(msg string) string
	report            ReportFunc
}

func buildOptions(tb testing.TB, opts []Option) *options {
	tb.Helper()
	o := &options{
		report: tb.Fatal,
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Option is an option for an assertion.
type Option func(*options)

// MessageTransform returns an [Option] that adds a message transform function.
// The function is called before the ReportFunc.
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

// Report returns an [Option] that sets the report function.
func Report(f ReportFunc) Option {
	return func(o *options) {
		o.report = f
	}
}
