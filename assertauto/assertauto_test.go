package assertauto

import (
	"testing"

	"github.com/pierrre/assert"
)

func TestEqual(t *testing.T) {
	ok := Equal(t, 123)
	assert.True(t, ok)
	ok = Equal(t, "foo")
	assert.True(t, ok)
}
