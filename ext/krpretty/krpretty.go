// Package krpretty provides an integration with github.com/kr/pretty.
package krpretty

import (
	"github.com/kr/pretty"
	"github.com/pierrre/assert"
)

// Configure configures the integration.
func Configure() {
	assert.ValueStringer = ValueStringer()
}

// ValueStringer returns the string representation of a value.
func ValueStringer() func(v any) string {
	return func(v any) string {
		return pretty.Sprint(v)
	}
}
