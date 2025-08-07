// Package assert provides utilities to assert conditions in tests.
//
// Assertion functions return a boolean value indicating whether the assertion succeeded.
//
// By default, assertion failures are reported using testing.TB.Fatal.
// It can be customized with the Report() option.
package assert

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pierrre/go-libs/bufpool"
	"github.com/pierrre/go-libs/runtimeutil"
	"github.com/pierrre/pretty"
)

// DefaultShowStack is the default value used to show stack traces on assertion failures, see [ShowStack] option.
var DefaultShowStack = true

// DefaultReport is the default [ReportFunc] used for assertion failures, see [Report] option.
var DefaultReport = testing.TB.Fatal

// ValueStringer is a function that returns the string representation of a value.
//
// It can be customized to provide a better string representation.
//
// By default it uses [pretty.String].
var ValueStringer func(any) string = pretty.String

// Fail handles assertion failure.
// It calls the [ReportFunc] with the given message.
func Fail(tb testing.TB, name string, msg string, stackSkip int, opts ...Option) {
	tb.Helper()
	msg = fmt.Sprintf("assert %s: %s", name, msg)
	o := buildOptions(opts)
	for _, f := range o.messageTransforms {
		msg = f(msg)
	}
	if o.showStack {
		buf := bufPool.Get()
		_, _ = buf.WriteString(msg)
		_, _ = buf.WriteString("\n\nStack trace:\n")
		for f := range runtimeutil.GetCallersFrames(runtimeutil.GetCallers(stackSkip + 1)) {
			if strings.HasPrefix(f.Function, "testing.") {
				break
			}
			_, _ = runtimeutil.WriteFrame(buf, f)
		}
		msg = buf.String()
		bufPool.Put(buf)
	}
	args := []any{msg}
	o.report(tb, args...)
}

var bufPool = &bufpool.Pool{
	MaxCap: -1,
}
