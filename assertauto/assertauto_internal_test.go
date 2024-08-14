package assertauto

import (
	"testing"
)

var DirectoryGlobal = directoryGlobal

func Update(u bool) Option {
	return update(u)
}

func InitTestFunction(tb testing.TB, dir string, name string, optfs ...Option) {
	tb.Helper()
	fp := getFilePath(dir, name)
	opts := buildOptions(optfs)
	getTestFunctionWithFile(tb, fp, opts)
}
