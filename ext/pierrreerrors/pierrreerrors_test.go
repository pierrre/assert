package pierrreerrors

import (
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
	"github.com/pierrre/errors"
)

func init() {
	Configure()
}

func Test(t *testing.T) {
	report := asserttest.NewReportPrefix(t, "assert no_error: error: error\nstack")
	ok := assert.NoError(t, errors.New("error"), assert.Report(report))
	assert.False(t, ok)
}
