package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestPointerNil(t *testing.T) {
	ok := PointerNil(t, (*int)(nil))
	True(t, ok)
}

func TestPointerNilFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert pointer_nil[int]: not nil")
	ok := PointerNil(t, new(int), Report(report))
	False(t, ok)
}

func TestPointerNotNil(t *testing.T) {
	ok := PointerNotNil(t, new(int))
	True(t, ok)
}

func TestPointerNotNilFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert pointer_not_nil[int]: nil")
	ok := PointerNotNil(t, (*int)(nil), Report(report))
	False(t, ok)
}
