package davecghspew

import (
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func init() {
	ConfigureDefault()
}

func Test(t *testing.T) {
	report := asserttest.NewReport(t, "assert equal[int]: not equal:\nv1 = (int) 1\n\nv2 = (int) 2\n")
	ok := assert.Equal(t, 1, 2, assert.Report(report))
	assert.False(t, ok)
}
