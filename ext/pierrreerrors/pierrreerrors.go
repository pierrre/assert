// Package pierrreerrors provides an integration with github.com/pierrre/errors.
package pierrreerrors

import (
	"github.com/pierrre/assert"
	"github.com/pierrre/errors/errverbose"
)

// Configure configures the integration.
//
// It sets assert.ErrorStringer to errverbose.String.
func Configure() {
	assert.ErrorStringer = errverbose.String //nolint:reassign // Replace the default implementation.
}
