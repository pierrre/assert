// Package assert provides utilities to assert conditions in tests.
//
// Assertion functions return a boolean value indicating whether the assertion succeeded.
//
// By default, assertion failures are reported using testing.TB.Fatal.
// This can be customized with the Report() option.
package assert

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pierrre/go-libs/bytesutil"
	"github.com/pierrre/go-libs/runtimeutil"
	"github.com/pierrre/pretty"
)

// DefaultShowStack is the default value used to show stack traces on assertion failures. See the [ShowStack] option.
var DefaultShowStack = true

// DefaultReport is the default [ReportFunc] used for assertion failures. See the [Report] option.
var DefaultReport = testing.TB.Fatal

// ValueStringer is a function that returns the string representation of a value.
//
// This can be customized to provide a better string representation.
//
// By default, it uses [pretty.String].
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
		bw := bytesWriterPool.Get()
		bw.AppendString(msg)
		bw.AppendString("\n\nStack trace:\n")
		for f := range runtimeutil.GetCallersFrames(runtimeutil.GetCallers(stackSkip + 1)) {
			if strings.HasPrefix(f.Function, "testing.") {
				break
			}
			*bw = runtimeutil.AppendFrame(*bw, f)
		}
		msg = bw.String()
		bytesWriterPool.Put(bw)
	}
	args := []any{msg}
	o.report(tb, args...)
}

var bytesWriterPool = &bytesutil.WriterPool{
	MaxCap: -1,
}
