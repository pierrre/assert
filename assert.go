// Package assert provides utilities to assert conditions in tests.
//
// Assertion functions return a boolean value indicating whether the assertion succeeded.
//
// By default, assertion failures are reported using testing.TB.Fatal.
// It can be customized with the Report() option.
package assert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pierrre/pretty"
)

// ReportFunc is a function that is called when an assertion fails.
//
// It is implemented by testing.TB.Fatal|Error|Skip|Log.
type ReportFunc func(args ...any)

// Fail handles assertion failure.
// It calls the ReportFunc with the given message.
func Fail(tb testing.TB, name string, msg string, opts ...Option) {
	tb.Helper()
	msg = fmt.Sprintf("assert %s: %s", name, msg)
	o := buildOptions(tb, opts)
	for _, f := range o.messageTransforms {
		msg = f(msg)
	}
	args := []any{msg}
	o.report(args...)
}

// ValueStringer is a function that returns the string representation of a value.
//
// It can be customized to provide a better string representation.
//
// By default it uses [pretty.String].
var ValueStringer = pretty.String

// TypeString returns a string representation of a type.
func TypeString[T any]() string {
	var v T
	// Use pointer in order to work with interface types.
	typ := reflect.TypeOf(&v).Elem()
	pkgPath := typ.PkgPath()
	if pkgPath != "" {
		return pkgPath + "." + typ.Name()
	}
	return typ.String()
}
