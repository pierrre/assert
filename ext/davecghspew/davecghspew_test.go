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
	report := asserttest.NewReportAuto(t)
	ok := assert.Equal(t, 1, 2, assert.Report(report))
	assert.False(t, ok)
}
