package assert_test

import (
	"io"
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/assertauto"
)

func TestTypeStringString(t *testing.T) {
	s := TypeString[string]()
	assertauto.Equal(t, s)
}

func TestTypeStringIOWriter(t *testing.T) {
	s := TypeString[io.Writer]()
	assertauto.Equal(t, s)
}
