// Package assertauto provides helpers to automatically update the expected values of test assertions.
//
// It compares actual values with expected values that are stored in files.
//
// Run the tests with the environment variable ASSERTAUTO_UPDATE=true to update the expected values.
//
// The values are stored in the "_assertauto" directory relative to the tested package.
// Each test creates a file with the name of the test and the ".txt" extension.
// Values of the same tests are stored sequentially in the same file.
//
// Values are converted to string using [DefaultValueStringer].
package assertauto

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/pierrre/assert"
	diff "github.com/pierrre/assert/assertauto/internal"
	"github.com/pierrre/go-libs/syncutil"
)

const (
	directoryEnvVar = "ASSERTAUTO_DIRECTORY"
	updateEnvVar    = "ASSERTAUTO_UPDATE"
)

var (
	// DefaultValueStringer is the default value stringer.
	DefaultValueStringer = assert.ValueStringer
	// DefaultDirectory is the default directory.
	DefaultDirectory = "_assertauto"
	// DefaultUpdate is the default update value.
	DefaultUpdate = false
)

func init() {
	initDefaultDirectory()
	initDefaultUpdate()
}

func initDefaultDirectory() {
	s, ok := os.LookupEnv(directoryEnvVar)
	if !ok || s == "" {
		return
	}
	DefaultDirectory = s
}

func initDefaultUpdate() {
	s, ok := os.LookupEnv(updateEnvVar)
	if !ok {
		return
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		err = fmt.Errorf("parse %s environment variable: %w", updateEnvVar, err)
		panic(err)
	}
	DefaultUpdate = b
}

var (
	initOnce  sync.Once
	initError error //nolint:errname // This is not a sentinel error.
)

func ensureInit() error {
	initOnce.Do(func() {
		initError = doInit()
	})
	return initError
}

func doInit() error {
	if DefaultUpdate {
		err := os.RemoveAll(DefaultDirectory)
		if err != nil {
			return fmt.Errorf("remove directory %q: %w", DefaultDirectory, err)
		}
	}
	return nil
}

func assertNoError(tb testing.TB, err error, opts *options) bool {
	tb.Helper()
	if err != nil {
		msg := err.Error() + "\n\nSee documentation at https://pkg.go.dev/github.com/pierrre/assert/assertauto\nRun the tests with the environment variable ASSERTAUTO_UPDATE=true to update the expected values."
		assert.Fail(tb, "assertauto", msg, opts.opts...)
		return false
	}
	return true
}

// Equal asserts that the value is equal to the expected value.
func Equal(tb testing.TB, v any, optfs ...Option) bool {
	tb.Helper()
	opts := buildOptions(tb.Name(), optfs)
	err := equal(tb, v, opts)
	return assertNoError(tb, err, opts)
}

func equal(tb testing.TB, v any, opts *options) error {
	tb.Helper()
	err := validateTestName(opts.testName)
	if err != nil {
		return fmt.Errorf("validate test name: %w", err)
	}
	err = ensureInit()
	if err != nil {
		return fmt.Errorf("init: %w", err)
	}
	s := opts.valueStringer(v)
	if strings.Contains(s, separator) {
		return errors.New("contains separator")
	}
	if opts.update {
		addValue(tb, s, opts)
	} else {
		expected, err := getValue(tb, opts)
		if err != nil {
			return fmt.Errorf("get value: %w", err)
		}
		if s != expected {
			d := diff.Diff("actual", []byte(s+"\n"), "expected", []byte(expected+"\n"))
			return fmt.Errorf("not equal:\n%s\nactual: %s\n\nexpected: %s", d, s, expected)
		}
	}
	return nil
}

func validateTestName(testName string) error {
	if strings.Contains(testName, "..") {
		return fmt.Errorf("contains \"..\": %q", testName)
	}
	return nil
}

// AllocsPerRun asserts that a function allocates a certain number of times per run.
func AllocsPerRun(tb testing.TB, runs int, f func(), optfs ...Option) (float64, bool) {
	tb.Helper()
	allocs := testing.AllocsPerRun(runs, f)
	v := allocsPerRun{
		Runs:   runs,
		Allocs: allocs,
	}
	return allocs, Equal(tb, v, optfs...)
}

type allocsPerRun struct {
	Runs   int
	Allocs float64
}

var values syncutil.Map[string, []string]

func addValue(tb testing.TB, v string, opts *options) {
	tb.Helper()
	vs, ok := values.Load(opts.testName)
	if !ok {
		setCleanupValues(tb, true, opts)
	}
	vs = append(vs, v)
	values.Store(opts.testName, vs)
}

func getValue(tb testing.TB, opts *options) (string, error) {
	tb.Helper()
	vs, ok := values.Load(opts.testName)
	if !ok {
		var err error
		vs, err = loadValues(opts)
		if err != nil {
			return "", fmt.Errorf("load values: %w", err)
		}
		setCleanupValues(tb, false, opts)
	}
	if len(vs) == 0 {
		return "", errors.New("no values left")
	}
	v := vs[0]
	vs = vs[1:]
	values.Store(opts.testName, vs)
	return v, nil
}

func setCleanupValues(tb testing.TB, save bool, opts *options) {
	tb.Helper()
	tb.Cleanup(func() {
		err := cleanupValues(tb, save, opts)
		assertNoError(tb, err, opts)
	})
}

func cleanupValues(tb testing.TB, save bool, opts *options) error {
	tb.Helper()
	vs, _ := values.LoadAndDelete(opts.testName)
	if !tb.Failed() {
		if save {
			return saveValues(vs, opts)
		}
		if len(vs) > 0 {
			return fmt.Errorf("remaining values:\n%s", encodeValues(vs))
		}
	}
	return nil
}

func saveValues(vs []string, opts *options) error {
	s := encodeValues(vs)
	fp := buildFilePath(opts.directory, opts.testName)
	dir := filepath.Dir(fp)
	err := os.MkdirAll(dir, 0o755) //nolint:gosec // We want 755.
	if err != nil {
		return fmt.Errorf("create directory: %w", err)
	}
	err = os.WriteFile(fp, []byte(s), 0o644) //nolint:gosec // We want 644.
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}

func encodeValues(vs []string) string {
	return strings.Join(vs, separator)
}

func loadValues(opts *options) ([]string, error) {
	fp := buildFilePath(opts.directory, opts.testName)
	b, err := os.ReadFile(fp) //nolint:gosec // We want to read the file.
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	return decodeValues(string(b)), nil
}

func decodeValues(s string) []string {
	return strings.Split(s, separator)
}

func buildFilePath(dir string, testName string) string {
	// TODO escape weird characters from test name ?
	return filepath.Join(dir, testName+".txt")
}

const separator = "\n\t========== assertauto ==========\n"

// Option is an option for assertauto.
type Option func(*options)

type options struct {
	valueStringer func(any) string
	directory     string
	testName      string
	update        bool
	opts          []assert.Option
}

func buildOptions(testName string, opts []Option) *options {
	o := newOptions(testName)
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func newOptions(testName string) *options {
	return &options{
		valueStringer: DefaultValueStringer,
		directory:     DefaultDirectory,
		testName:      testName,
		update:        DefaultUpdate,
	}
}

// ValueStringer returns an [Option] that sets the value stringer.
func ValueStringer(vs func(any) string) Option {
	return func(o *options) {
		o.valueStringer = vs
	}
}

func directory(d string) Option {
	return func(o *options) {
		o.directory = d
	}
}

func testName(tn string) Option {
	return func(o *options) {
		o.testName = tn
	}
}

func update(u bool) Option {
	return func(o *options) {
		o.update = u
	}
}

// AssertOptions returns an [Option] that sets the assert options.
func AssertOptions(opts ...assert.Option) Option {
	return func(o *options) {
		o.opts = opts
	}
}
