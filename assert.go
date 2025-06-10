// Package assert provides utilities to assert conditions in tests.
//
// Assertion functions return a boolean value indicating whether the assertion succeeded.
//
// By default, assertion failures are reported using testing.TB.Fatal.
// It can be customized with the Report() option.
package assert

import (
	"fmt"
	"testing"

	"github.com/pierrre/pretty"
)

// Fail handles assertion failure.
// It calls the [ReportFunc] with the given message.
func Fail(tb testing.TB, name string, msg string, opts ...Option) {
	tb.Helper()
	msg = fmt.Sprintf("assert %s: %s", name, msg)
	o := buildOptions(opts)
	for _, f := range o.messageTransforms {
		msg = f(msg)
	}
	args := []any{msg}
	o.report(tb, args...)
}

// ValueStringer is a function that returns the string representation of a value.
//
// It can be customized to provide a better string representation.
//
// By default it uses [pretty.String].
var ValueStringer func(any) string = pretty.String
