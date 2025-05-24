package assertauto_test

import (
	"os"
	"testing"

	"github.com/pierrre/assert"
	. "github.com/pierrre/assert/assertauto"
	"github.com/pierrre/assert/asserttest"
)

func Test(t *testing.T) {
	Equal(t, nil)
	Equal(t, true)
	Equal(t, 123456)
	Equal(t, 123.456)
	Equal(t, "test")
}

func TestEqual(t *testing.T) {
	tmpDir := t.TempDir()
	testName := t.Name() + "Fake"
	t.Run("Save", func(t *testing.T) {
		ok := Equal(t, "test", Directory(tmpDir), TestName(testName), Update(true))
		assert.True(t, ok)
	})
	t.Run("Load", func(t *testing.T) {
		ok := Equal(t, "test", Directory(tmpDir), TestName(testName), Update(false))
		assert.True(t, ok)
	})
}

func TestEqualFail(t *testing.T) {
	tmpDir := t.TempDir()
	testName := t.Name() + "Fake"
	t.Run("Save", func(t *testing.T) {
		ok := Equal(t, "aaa", Directory(tmpDir), TestName(testName), Update(true))
		assert.True(t, ok)
	})
	report := asserttest.NewReportAuto(t)
	t.Run("Load", func(t *testing.T) {
		ok := Equal(t, "bbb", Directory(tmpDir), TestName(testName), Update(false), AssertOptions(assert.Report(report)))
		assert.False(t, ok)
	})
}

func TestAllocsPerRun(t *testing.T) {
	allocs, ok := AllocsPerRun(t, 100, func() {
		_ = make([]byte, 1<<20)
	})
	assert.True(t, ok)
	Equal(t, allocs)
}

func TestErrorTestName(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	t.Run("../aaa", func(t *testing.T) {
		Equal(t, "test", AssertOptions(assert.Report(report)))
	})
}

func TestErrorContainsSeparator(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := Equal(t, Separator, AssertOptions(assert.Report(report)), ValueStringer(func(v any) string {
		return "test" + Separator + "test"
	}))
	assert.False(t, ok)
}

func TestErrorNoValuesLeft(t *testing.T) {
	tmpDir := t.TempDir()
	testName := t.Name() + "Fake"
	t.Run("Save", func(t *testing.T) {
		ok := Equal(t, "test", Directory(tmpDir), TestName(testName), Update(true))
		assert.True(t, ok)
	})
	report := asserttest.NewReportAuto(t)
	t.Run("Load", func(t *testing.T) {
		ok := Equal(t, "test", Directory(tmpDir), TestName(testName), Update(false))
		assert.True(t, ok)
		ok = Equal(t, "test", Directory(tmpDir), TestName(testName), Update(false), AssertOptions(assert.Report(report)))
		assert.False(t, ok)
	})
}

func TestErrorRemainingValues(t *testing.T) {
	tmpDir := t.TempDir()
	testName := t.Name() + "Fake"
	t.Run("Save", func(t *testing.T) {
		Equal(t, "test", Directory(tmpDir), TestName(testName), Update(true))
		Equal(t, "test", Directory(tmpDir), TestName(testName), Update(true))
	})
	report := asserttest.NewReportAuto(t)
	t.Run("Load", func(t *testing.T) {
		ok := Equal(t, "test", Directory(tmpDir), TestName(testName), Update(false), AssertOptions(assert.Report(report)))
		assert.True(t, ok)
	})
}

func TestErrorCreateDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	testName := t.Name() + "Fake"
	fp := BuildFilePath(tmpDir, testName)
	err := os.WriteFile(fp, []byte("test"), 0o644) //nolint:gosec // We want 644.
	assert.NoError(t, err)
	report := asserttest.NewReportPrefix(t, "assert assertauto: create directory:")
	t.Run("Save", func(t *testing.T) {
		ok := Equal(t, "test", Directory(fp), Update(true), AssertOptions(assert.Report(report)))
		assert.True(t, ok)
	})
}

func TestErrorWriteFile(t *testing.T) {
	tmpDir := t.TempDir()
	testName := t.Name() + "Fake"
	fp := BuildFilePath(tmpDir, testName)
	err := os.MkdirAll(fp, 0o750)
	assert.NoError(t, err)
	report := asserttest.NewReportPrefix(t, "assert assertauto: write file:")
	t.Run("Save", func(t *testing.T) {
		ok := Equal(t, "test", Directory(tmpDir), TestName(testName), Update(true), AssertOptions(assert.Report(report)))
		assert.True(t, ok)
	})
}

func TestErrorReadFile(t *testing.T) {
	tmpDir := t.TempDir()
	testName := t.Name() + "Fake"
	ok := Equal(t, "test", Directory(tmpDir), TestName(testName), Update(false), AssertOptions(assert.Report(func(_ testing.TB, args ...any) {})))
	assert.False(t, ok)
}
