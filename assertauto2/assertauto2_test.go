package assertauto2_test

import (
	"testing"

	. "github.com/pierrre/assert/assertauto2"
)

func TestFoo(t *testing.T) {
	AllocsPerRun(t, 100, func() {
		_ = make([]byte, 1<<20)
	})
}
