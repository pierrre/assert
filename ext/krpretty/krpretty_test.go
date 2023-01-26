package krpretty

import (
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func init() {
	Configure()
}

func Test(t *testing.T) {
	report := asserttest.NewReport(t, "assert equal[int]: not equal:\nv1 = int(1)\nv2 = int(2)")
	ok := assert.Equal(t, 1, 2, assert.Report(report))
	assert.False(t, ok)
}
