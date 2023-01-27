# Assert

Go test assertion library.

[![Go Reference](https://pkg.go.dev/badge/github.com/pierrre/assert.svg)](https://pkg.go.dev/github.com/pierrre/assert)

## Features

- [Test assertion (equal, comparison, nil, empty, length, error, etc...)](#assertions)
- [No reflection (uses generics)](#why-)
- [Third party integration (print value, value comparison)](#integrations)

## Assertions

A simple assertion:

```go
assert.Equal(t, value, 1)
```

By default, assertions fail with `Fatal()`.
It can be changed with the `Report()` option:

```go
assert.Equal(t, value, 1, assert.Report(t.Error))
```

The report message can be customized:

```go
assert.Equal(t, value, 1, assert.MessageWrap("test"))
```

## Why ?

This assertion library is an experiment to see if it is possible to do better than `github.com/stretchr/testify` with the new features offered by the recent Go versions.

```go
func Test(t *testing.T) {
    value := getValue()
    require.Equal(t, 1, value)
}

func getValue() int64 {
    return 1
}
```

Surprinsingly, this test fails with this error:

```text
Error: Not equal:
expected: int(1)
actual  : int64(1)
```

This issue is caused by the types, which are no identical (the `1` constant is an `int` and not an `int64`), and it's possible to fix it:

Convert the value to `int64`:

```go
require.Equal(t, int64(1), value)
```

Use `EqualValues()` which converts the values to the same type:

```go
require.EqualValues(t, 1, value)
```

But the internal implementation is not simple: it requires heavy usage of reflection, and the code is quite complex.
What if we could simply use the `==` operator ?

This is the solution chosen by this library.
It uses generics to do the comparison, and it works with any types:

```go
assert.Equal(t, 1, value)
```

The constant `1` is automatically converted to the type of the `value` variable without using reflection.

However, this approchach has a limitation: it requires to write a different assertion function for each "kind" (map, slice, etc...)

## Integrations

Third party integrations allow to customize how values are printed and compared:

- [`davecghspew`](https://pkg.go.dev/github.com/pierrre/assert/ext/davecghspew) prints values with [`github.com/davecgh/go-spew`](https://pkg.go.dev/github.com/davecgh/go-spew/spew)
- [`krpretty`](https://pkg.go.dev/github.com/pierrre/assert/ext/krpretty) prints values with [`github.com/kr/pretty`](https://pkg.go.dev/github.com/kr/pretty)
- [`pierrrecompare`](https://pkg.go.dev/github.com/pierrre/assert/ext/pierrrecompare) compares values with [`github.com/pierrre/compare`](https://pkg.go.dev/github.com/pierrre/compare)
- [`googlecmp`](https://pkg.go.dev/github.com/pierrre/assert/ext/googlecmp) compares values with [`github.com/google/go-cmp`](https://pkg.go.dev/github.com/google/go-cmp/cmp)
- [`pierrreerrors`](https://pkg.go.dev/github.com/pierrre/assert/ext/pierrreerrors) prints errors with [`github.com/pierrre/errors`](https://pkg.go.dev/github.com/pierrre/errors)

## FAQ

> Why not use `github.com/stretchr/testify` ?

I think it's a great library, but I wanted to [try something different](#why-).
I also wanted to try generics, and to see if it was possible to make an assertion library without reflection.

> Where are `Nil()` and `NotNil()` ?

- For slices use [SliceNil()](https://pkg.go.dev/github.com/pierrre/assert#SliceNil) and [SliceNotNil()](https://pkg.go.dev/github.com/pierrre/assert#SliceNotNil)
- For maps use [MapNil()](https://pkg.go.dev/github.com/pierrre/assert#MapNil) and [MapNotNil()](https://pkg.go.dev/github.com/pierrre/assert#MapNotNil)
- For comparable types use [Zero()](https://pkg.go.dev/github.com/pierrre/assert#Zero) and [NotZero()](https://pkg.go.dev/github.com/pierrre/assert#NotZero)
- For interface types use [Nil()](https://pkg.go.dev/github.com/pierrre/assert#Nil) and [NotNil()](https://pkg.go.dev/github.com/pierrre/assert#NotNil), but they will be removed once Go 1.20 is supported.
