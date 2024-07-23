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

func TestTypeStringAllocsString(t *testing.T) {
	AllocsPerRun(t, 100, func() {
		TypeString[string]()
	}, 0)
}

func TestTypeStringIOWriter(t *testing.T) {
	s := TypeString[io.Writer]()
	assertauto.Equal(t, s)
}

func TestTypeStringAllocsIOWriter(t *testing.T) {
	AllocsPerRun(t, 100, func() {
		TypeString[io.Writer]()
	}, 0)
}
