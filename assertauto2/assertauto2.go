package assertauto2

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
	"github.com/pierrre/pretty"
)

const (
	directoryEnvVar  = "ASSERTAUTO2_DIRECTORY"
	directoryDefault = "_assertauto2"
	updateEnvVar     = "ASSERTAUTO_UPDATE"
	updateDefault    = false
)

var (
	ValueStringer   = assert.ValueStringer
	directoryGlobal = initDirectoryGlobal()
	updateGlobal    = initUpdateGlobal()
)

func init() {
	pretty.DefaultCommonValueWriter.ConfigureTest()
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

// Equal asserts that the value is equal to the expected value.
func Equal(tb testing.TB, v any, optfs ...Option) bool {
	tb.Helper()
	opts := buildOptions(optfs)
	err := equal(tb, v, opts)
	return assert.NoError(tb, err, opts.opts...) // TODO improve message here
}

func equal(tb testing.TB, v any, opts *options) error {
	tb.Helper()
	s := ValueStringer(v)
	if opts.update {
		addValue(tb, s, opts)
	} else {
		v, err := getValue(tb, opts)
		if err != nil {
			return fmt.Errorf("get entry: %w", err)
		}
		if v != s {
			return errors.New("not equal") // TODO improve diff
		}
	}
	return nil
}

var values = &syncutil.Map[string, []string]{}

func addValue(tb testing.TB, e string, opts *options) {
	tb.Helper()
	testName := tb.Name()
	vs, ok := values.Load(testName)
	if !ok {
		tb.Cleanup(func() {
			err := cleanupValues(testName, true, opts)
			assert.NoError(tb, err, opts.opts...) // TODO improve message here
		})
	}
	vs = append(vs, e)
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
		tb.Cleanup(func() {
			err := cleanupValues(testName, false, opts)
			assert.NoError(tb, err, opts.opts...) // TODO improve message here
		})
	}
	if len(vs) == 0 {
		return "", errors.New("no values left")
	}
	v := vs[0]
	vs = vs[1:]
	values.Store(testName, vs)
	return v, nil
}

func cleanupValues(testName string, save bool, opts *options) error {
	vs, _ := values.LoadAndDelete(testName)
	if save {
		return saveValues(testName, vs, opts)
	}
	if len(vs) > 0 {
		return errors.New("values left") // TODO show values left in error
	}
	return nil
}

func saveValues(testName string, vs []string, opts *options) error {
	s := strings.Join(vs, separator)
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

func loadValues(testName string, opts *options) ([]string, error) {
	fp := getFilePath(testName, opts)
	b, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	return strings.Split(string(b), separator), nil
}

// // AllocsPerRun asserts that a function allocates a certain number of times per run.
// func AllocsPerRun(tb testing.TB, runs int, f func(), optfs ...Option) bool {
// 	tb.Helper()
// 	allocs := testing.AllocsPerRun(runs, f)
// 	v := allocsPerRun{
// 		Runs:   runs,
// 		Allocs: allocs,
// 	}
// 	return Equal(tb, v, optfs...)
// }

// type allocsPerRun struct {
// 	Runs   int
// 	Allocs float64
// }

func getFilePath(testName string, opts *options) string {
	// TODO escape test name ?
	return filepath.Join(opts.directory, testName+".txt")
}

const separator = "\n==========================================\n"

type options struct {
	directory string
	name      string
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

// Name returns an [Option] that sets the name of the entry.
func Name(name string) Option {
	return func(o *options) {
		o.name = name
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
