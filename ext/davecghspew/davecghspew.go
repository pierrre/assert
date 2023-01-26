// Package davecghspew provides an integration with github.com/davecgh/go-spew.
package davecghspew

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/pierrre/assert"
)

// Configure configures the integration.
func Configure(config *spew.ConfigState) {
	assert.ValueStringer = NewValueStringer(config)
}

// NewValueStringer returns a function that returns the string representation of a value.
func NewValueStringer(config *spew.ConfigState) func(v any) string {
	return func(v any) string {
		return config.Sdump(v)
	}
}
