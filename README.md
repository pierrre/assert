# Assert

Go test assertion library.

[![Go Reference](https://pkg.go.dev/badge/github.com/pierrre/assert.svg)](https://pkg.go.dev/github.com/pierrre/assert)

## Features

- [Test assertion (equal, comparison, nil, empty, length, error, etc...)](#assertions)
- [No reflection (uses generics)](#why-)
- [Customization (print and compare values)](#customization)

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

This assertion library is an experiment to see if it is possible to do better than `github.com/stretchr/testify`, by using generics.

Here is an example of an issue with `github.com/stretchr/testify`:

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

But the internal implementation is not simple: it requires [heavy usage of reflection](https://github.com/stretchr/testify/blob/master/assert/assertion_compare.go), and the code is [quite complex](https://github.com/stretchr/testify/blob/master/assert/assertions.go).

What if we could simply use the `==` operator ?
This is the solution chosen by this library.
It uses generics to do the comparison, and it works with any comparable type:

```go
func Equal[T comparable](tb testing.TB, v1, v2 T, opts ...Option) bool {
    tb.Helper()
    ok := v1 == v2
    if !ok {
        Fail(...)
    }
    return ok
}
```

```go
assert.Equal(t, 1, value)
```

The constant `1` is automatically converted to the type of the `value` variable without using reflection.

However, this approchach has a limitation: it requires to write a different assertion function for each "kind" (map, slice, etc...)

## Customization

The default behavior can be customized:

- [`DeepEqualer`](https://pkg.go.dev/github.com/pierrre/assert#DeepEqualer) allows to customize how values are compared with [`DeepEqual()`](https://pkg.go.dev/github.com/pierrre/assert#DeepEqual).
- [`ValueStringer`](https://pkg.go.dev/github.com/pierrre/assert#ValueStringer) allows to customize how values are printed.
- [`ErrorStringer`](https://pkg.go.dev/github.com/pierrre/assert#ErrorStringer) allows to customize how errors are printed.

## FAQ

### Why not use `github.com/stretchr/testify` ?

I think it's a great library, but I wanted to [try something different](#why-).
I also wanted to try generics, and to see if it was possible to make an assertion library without reflection.

### Where are `Nil()` and `NotNil()` ?

- For slices use [SliceNil()](https://pkg.go.dev/github.com/pierrre/assert#SliceNil) and [SliceNotNil()](https://pkg.go.dev/github.com/pierrre/assert#SliceNotNil)
- For maps use [MapNil()](https://pkg.go.dev/github.com/pierrre/assert#MapNil) and [MapNotNil()](https://pkg.go.dev/github.com/pierrre/assert#MapNotNil)
- For comparable types use [Zero()](https://pkg.go.dev/github.com/pierrre/assert#Zero) and [NotZero()](https://pkg.go.dev/github.com/pierrre/assert#NotZero)
