// Package assert provides utilities to assert conditions in tests.
//
// Assertion functions return a boolean value indicating whether the assertion succeeded.
//
// By default, assertion failures are reported using testing.TB.Fatal.
// This can be customized with the Report() option.
package assert

import (
	"fmt"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/pierrre/go-libs/bytesutil"
	"github.com/pierrre/go-libs/runtimeutil"
	"github.com/pierrre/go-libs/syncutil/atomicutil"
	"github.com/pierrre/pretty"
)

// DefaultShowStack is the default value used to show stack traces on assertion failures. See the [ShowStack] option.
//
// By default it is true.
var DefaultShowStack atomic.Bool

func init() {
	DefaultShowStack.Store(true)
}

// DefaultReport is the default [ReportFunc] used for assertion failures. See the [Report] option.
//
// By default it uses [testing.TB.Fatal].
var DefaultReport atomicutil.Value[ReportFunc]

func init() {
	DefaultReport.Store(ReportFunc(testing.TB.Fatal))
}

// ValueStringer is a function that returns the string representation of a value.
//
// By default, it uses [pretty.String].
var ValueStringer atomicutil.Value[func(any) string]

func init() {
	ValueStringer.Store(pretty.String)
}

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
		defer bytesWriterPool.Put(bw)
		bw.AppendString(msg)
		bw.AppendString("\n\nStack trace:\n")
		fs := runtimeutil.GetCallersFrames(runtimeutil.GetCallers(stackSkip + 1))
		fs(func(f runtime.Frame) bool {
			if strings.HasPrefix(f.Function, "testing.") {
				return false
			}
			*bw = runtimeutil.AppendFrame(*bw, f)
			return true
		})
		msg = bw.String()
	}
	args := []any{msg}
	o.report(tb, args...)
}

var bytesWriterPool = &bytesutil.WriterPool{
	MaxCap: -1,
}
