package assertauto_test

import (
	"testing"

	"github.com/pierrre/assert"
	. "github.com/pierrre/assert/assertauto"
	"github.com/pierrre/assert/asserttest"
	"github.com/pierrre/pretty"
)

func init() {
	pretty.DefaultCommonValueWriter.ConfigureTest()
}

func TestEqual(t *testing.T) {
	test(t, func(t *testing.T, opts ...Option) { //nolint:thelper // This is not a helper.
		ok := Equal(t, 123, append(opts, Name("int"))...)
		assert.True(t, ok)
		ok = Equal(t, "foo", append(opts, Name("string"))...)
		assert.True(t, ok)
	})
}

func TestEqualFail(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		ok := Equal(t, 123, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		ok := Equal(t, 456, Update(false), Directory(tmpDir), FileName(rootTest.Name()), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestEqualFailName(t *testing.T) {
	testFailName(t, func(t *testing.T, opts ...Option) bool { //nolint:thelper // This is not a helper.
		return Equal(t, 123, opts...)
	})
}

func TestEqualFailEntryType(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		ok := DeepEqual(t, &testSTruct{
			Foo: "bar",
		}, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		ok := Equal(t, 123, Update(false), Directory(tmpDir), FileName(rootTest.Name()), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestDeepEqual(t *testing.T) {
	test(t, func(t *testing.T, opts ...Option) { //nolint:thelper // This is not a helper.
		ok := DeepEqual(t, &testSTruct{
			Foo: "bar",
		}, opts...)
		assert.True(t, ok)
	})
}

func TestDeepEqualFail(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		ok := DeepEqual(t, &testSTruct{
			Foo: "bar",
		}, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		ok := DeepEqual(t, &testSTruct{
			Foo: "baz",
		}, Update(false), Directory(tmpDir), FileName(rootTest.Name()), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestDeepEqualFailName(t *testing.T) {
	testFailName(t, func(t *testing.T, opts ...Option) bool { //nolint:thelper // This is not a helper.
		return DeepEqual(t, &testSTruct{
			Foo: "bar",
		}, opts...)
	})
}

func TestDeepEqualFailEntryType(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		ok := Equal(t, 123, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		ok := DeepEqual(t, &testSTruct{
			Foo: "bar",
		}, Update(false), Directory(tmpDir), FileName(rootTest.Name()), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestAllocsPerRun(t *testing.T) {
	test(t, func(t *testing.T, opts ...Option) { //nolint:thelper // This is not a helper.
		ok := AllocsPerRun(t, 10, func() {
			_ = make([]byte, 1<<20)
		}, opts...)
		assert.True(t, ok)
	})
}

func TestAllocsPerRunFail(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		ok := AllocsPerRun(t, 10, func() {}, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		ok := AllocsPerRun(t, 10, func() {
			_ = make([]byte, 1<<20)
		}, Update(false), Directory(tmpDir), FileName(rootTest.Name()), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestAllocsPerRunFailName(t *testing.T) {
	testFailName(t, func(t *testing.T, opts ...Option) bool { //nolint:thelper // This is not a helper.
		return AllocsPerRun(t, 10, func() {
			_ = make([]byte, 1<<20)
		}, opts...)
	})
}

func TestAllocsPerRunFailEntryType(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		ok := Equal(t, 123, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		ok := AllocsPerRun(t, 10, func() {
			_ = make([]byte, 1<<20)
		}, Update(false), Directory(tmpDir), FileName(rootTest.Name()), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestAllocsPerRunFailRuns(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		ok := AllocsPerRun(t, 10, func() {
			_ = make([]byte, 1<<20)
		}, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		ok := AllocsPerRun(t, 20, func() {
			_ = make([]byte, 1<<20)
		}, Update(false), Directory(tmpDir), FileName(rootTest.Name()), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestFailEntryNoRemaining(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		ok := Equal(t, 123, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		ok := Equal(t, 123, Update(false), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
		ok = Equal(t, 123, Update(false), Directory(tmpDir), FileName(rootTest.Name()), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestFailEntriesRemaining(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		ok := Equal(t, 123, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
		ok = Equal(t, 123, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		ok := Equal(t, 123, Update(false), Directory(tmpDir), FileName(rootTest.Name()), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.True(t, ok)
	})
}

func test(t *testing.T, f func(t *testing.T, opts ...Option)) {
	t.Helper()
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		f(t, Update(true), Directory(tmpDir), FileName(rootTest.Name()))
	})
	t.Run("Read", func(t *testing.T) {
		f(t, Update(false), Directory(tmpDir), FileName(rootTest.Name()))
	})
	t.Run("Normal", func(t *testing.T) {
		f(t, FileName(rootTest.Name()))
	})
}

func testFailName(t *testing.T, f func(t *testing.T, opts ...Option) bool) {
	t.Helper()
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		ok := f(t, Update(true), Directory(tmpDir), FileName(rootTest.Name()), Name("foo"))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		ok := f(t, Update(false), Directory(tmpDir), FileName(rootTest.Name()), Name("bar"), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

type testSTruct struct {
	Foo string
}
