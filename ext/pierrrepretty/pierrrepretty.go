// Package pierrrepretty provides an integration with github.com/pierrre/pretty.
package pierrrepretty

import (
	"github.com/pierrre/assert"
	"github.com/pierrre/pretty"
)

// Configure configures the integration.
//
// It sets assert.ValueStringer with pretty.Config.String.
func Configure(config *pretty.Config) {
	assert.ValueStringer = config.String
}

// ConfigureDefault calls Configure() with pretty.DefaultConfig.
func ConfigureDefault() {
	Configure(pretty.DefaultConfig)
}
