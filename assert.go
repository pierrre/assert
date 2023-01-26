// Package assert provides utilities to assert conditions in tests.
package assert

import (
	"fmt"
	"reflect" //nolint:depguard // Used for type string representation.
	"strconv"
	"testing"
)

// ReportFunc is a function that is called when an assertion fails.
//
// It is implemented by testing.TB.Fatal|Error|Skip|Log.
//
// By default it is Fatal.
// It can be changed with Report().
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

// ValueStringer is a function that returns a string representation of a value.
//
// It can be customized to provide a better string representation.
//
// By default it uses fmt.Sprint, or strconv.Quote for strings.
var ValueStringer = func(v any) string {
	if v, ok := v.(string); ok {
		return strconv.Quote(v)
	}
	return fmt.Sprint(v)
}

// TypeString returns a string representation of a type.
func TypeString[T any]() string {
	var v T
	// Use pointer in order to work with interface types.
	return reflect.TypeOf(&v).Elem().String()
}
