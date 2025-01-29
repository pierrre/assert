package assertauto2

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"

	"github.com/pierrre/assert"
)

const (
	directoryEnvVar  = "ASSERTAUTO2_DIRECTORY"
	directoryDefault = "_assertauto2"
	updateEnvVar     = "ASSERTAUTO_UPDATE"
	updateDefault    = false
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

// Equal asserts that the value is equal to the expected value.
func Equal(tb testing.TB, v any, optfs ...Option) bool {
	tb.Helper()
	opts := buildOptions(optfs)
	s := assert.ValueStringer(v)
	fp := getFilePath(tb)
	if opts.update {
		dir := filepath.Dir(fp)
		err := os.MkdirAll(dir, 0o755)
		assert.NoError(tb, err)
		err = os.WriteFile(fp, []byte(s), 0o644) //nolint:gosec // We want 644.
		assert.NoError(tb, err)
		return true
	}
	expected, err := os.ReadFile(fp)
	assert.NoError(tb, err)
	if s != string(expected) {
		assert.Fail(tb, "assertauto equal", fmt.Sprintf("not equal:\ngot %s\nexpected %s", s, string(expected)), opts.opts...)
		return false
	}
	return true
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

func getFilePath(tb testing.TB) string {
	tb.Helper()
	return filepath.Join(directoryGlobal, tb.Name()+".txt")
}

type testFunction struct {
	mu      sync.Mutex
	entries [][]byte
}

const separatorFormat = "========== %s =========="

func getSeparator(tb testing.TB) string {
	tb.Helper()
	h := sha256.Sum256([]byte(tb.Name()))
	return fmt.Sprintf(separatorFormat, string(h[:]))
}

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
