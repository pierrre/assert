// Package assertauto provides helpers to automatically update the expected values of assertions.
//
// This is highly experimental an not yet ready for public usage.
package assertauto

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/go-libs/syncutil"
	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	directoryEnvVar  = "ASSERTAUTO_DIRECTORY"
	directoryDefault = "_assertauto"
	updateEnvVar     = "ASSERTAUTO_UPDATE"
	updateDefault    = false
)

var (
	ValueStringer  = assert.ValueStringer
	DiffMatchPatch = diffmatchpatch.New()
)

var (
	directoryGlobal = initDirectoryGlobal()
	updateGlobal    = initUpdateGlobal()
)

func init() {
	if updateGlobal {
		err := os.RemoveAll(directoryGlobal)
		if err != nil {
			panic(err)
		}
	}
}

func initDirectoryGlobal() string {
	s, ok := os.LookupEnv(directoryEnvVar)
	if !ok {
		return directoryDefault
	}
	if s == "" {
		return directoryDefault
	}
	return s
}

func initUpdateGlobal() bool {
	s, ok := os.LookupEnv(updateEnvVar)
	if !ok {
		return updateDefault
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		err = fmt.Errorf("parse %s environment variable: %w", updateEnvVar, err)
		panic(err)
	}
	return b
}

func assertNoError(tb testing.TB, err error, opts *options) bool {
	tb.Helper()
	if err != nil {
		assert.Fail(tb, "assertauto", err.Error(), opts.opts...)
		return false
	}
	return true
}

// Equal asserts that the value is equal to the expected value.
func Equal(tb testing.TB, v any, optfs ...Option) bool {
	tb.Helper()
	opts := buildOptions(optfs)
	err := equal(tb, v, opts)
	return assertNoError(tb, err, opts)
}

func equal(tb testing.TB, v any, opts *options) error {
	tb.Helper()
	s := ValueStringer(v)
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
			diffs := DiffMatchPatch.DiffMain(expected, s, false)
			diffText := DiffMatchPatch.DiffPrettyText(diffs)
			return fmt.Errorf("not equal: %s\nactual: %s\n\nexpected: %s", diffText, s, expected)
		}
	}
	return nil
}

// AllocsPerRun asserts that a function allocates a certain number of times per run.
func AllocsPerRun(tb testing.TB, runs int, f func(), optfs ...Option) bool {
	tb.Helper()
	allocs := testing.AllocsPerRun(runs, f)
	v := allocsPerRun{
		Runs:   runs,
		Allocs: allocs,
	}
	return Equal(tb, v, optfs...)
}

type allocsPerRun struct {
	Runs   int
	Allocs float64
}

var values = &syncutil.Map[string, []string]{}

func addValue(tb testing.TB, v string, opts *options) {
	tb.Helper()
	testName := tb.Name()
	vs, ok := values.Load(testName)
	if !ok {
		setCleanupValues(tb, true, opts)
	}
	vs = append(vs, v)
	values.Store(testName, vs)
}

func getValue(tb testing.TB, opts *options) (string, error) {
	tb.Helper()
	testName := tb.Name()
	vs, ok := values.Load(testName)
	if !ok {
		var err error
		vs, err = loadValues(testName, opts)
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
	values.Store(testName, vs)
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
	testName := tb.Name()
	vs, _ := values.LoadAndDelete(testName)
	if tb.Failed() {
		return nil
	}
	if save {
		return saveValues(testName, vs, opts)
	}
	if len(vs) > 0 {
		return fmt.Errorf("remaining values:\n%s", encodeValues(vs))
	}
	return nil
}

func saveValues(testName string, vs []string, opts *options) error {
	s := encodeValues(vs)
	fp := getFilePath(testName, opts)
	dir := filepath.Dir(fp)
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return fmt.Errorf("mkdir all: %w", err)
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

func loadValues(testName string, opts *options) ([]string, error) {
	fp := getFilePath(testName, opts)
	b, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	return decodeValues(string(b)), nil
}

func decodeValues(s string) []string {
	return strings.Split(s, separator)
}

func getFilePath(testName string, opts *options) string {
	// TODO escape test name ?
	return filepath.Join(opts.directory, testName+".txt")
}

const separator = "\n\t========== assertauto ==========\n"

type options struct {
	directory string
	update    bool
	opts      []assert.Option
}

func buildOptions(opts []Option) *options {
	o := newOptions()
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func newOptions() *options {
	return &options{
		directory: directoryGlobal,
		update:    updateGlobal,
	}
}

func directory(dir string) Option {
	return func(o *options) {
		o.directory = dir
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

// Option is an option for assertauto.
type Option func(*options)
