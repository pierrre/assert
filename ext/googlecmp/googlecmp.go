// Package googlecmp provides an integration with github.com/google/go-cmp.
package googlecmp

import (
	"github.com/google/go-cmp/cmp"
	"github.com/pierrre/assert"
)

// Configure configures the integration.
//
// It sets assert.DeepEqualer with the result of NewDeepEqualer().
func Configure(opts ...cmp.Option) {
	assert.DeepEqualer = NewDeepEqualer(opts...)
}

// NewDeepEqualer returns a function that performs a deep equal comparison between 2 values.
func NewDeepEqualer(opts ...cmp.Option) func(v1, v2 any) (diff string, equal bool) {
	return func(v1, v2 any) (diff string, equal bool) {
		diff = cmp.Diff(v1, v2, opts...)
		if diff != "" {
			return diff, false
		}
		return "", true
	}
}
