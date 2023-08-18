// Package pierrrecompare provides an integration with github.com/pierrre/compare.
package pierrrecompare

import (
	"fmt"

	"github.com/pierrre/assert"
	"github.com/pierrre/compare"
)

// Configure configures the integration.
//
// It sets assert.DeepEqualer with DeepEqualer.
func Configure() {
	assert.DeepEqualer = DeepEqualer
}

// DeepEqualer performs a deep equal comparison betweetn v1 and v2.
func DeepEqualer(v1, v2 any) (diff string, equal bool) {
	res := compare.Compare(v1, v2)
	if len(res) == 0 {
		return "", true
	}
	diff = fmt.Sprintf("%+v", res)
	return diff, false
}
