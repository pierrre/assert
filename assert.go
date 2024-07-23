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
	"sync"
	"testing"

	"github.com/pierrre/pretty"
)

// ReportFunc is a function that is called when an assertion fails.
//
// It is implemented by [testing.TB.Fatal]|[testing.TB.Error]|[testing.TB.Skip]|[testing.TB.Log].
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
var ValueStringer func(any) string = pretty.String

var (
	typeStringCache   = make(map[reflect.Type]string)
	typeStringCacheMu sync.Mutex
)

// TypeString returns a string representation of a type.
func TypeString[T any]() string {
	typ := reflect.TypeFor[T]()
	typeStringCacheMu.Lock()
	defer typeStringCacheMu.Unlock()
	name, ok := typeStringCache[typ]
	if !ok {
		pkgPath := typ.PkgPath()
		if pkgPath != "" {
			name = pkgPath + "." + typ.Name()
		} else {
			name = typ.String()
		}
		typeStringCache[typ] = name
	}
	return name
}
