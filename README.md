must
====

[![Build](https://github.com/akshayjshah/must/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/akshayjshah/must/actions/workflows/ci.yaml)
[![Report Card](https://goreportcard.com/badge/go.akshayshah.org/must)](https://goreportcard.com/report/go.akshayshah.org/must)
[![GoDoc](https://pkg.go.dev/badge/go.akshayshah.org/must.svg)](https://pkg.go.dev/go.akshayshah.org/must)

`must` provides a generic equivalent of the Go standard library's
[template.Must][] or [regexp.MustCompile][]: a wrapper that converts errors to
panics. This is usually undesirable, but it's both convenient and safe when
creating global variables from pure functions and static inputs.

## Installation

```
go get go.akshayshah.org/must
```

## Usage

```go
package main

import (
	"fmt"
	"net/url"
)

var google = Must(url.Parse("https://www.google.com"))

func main() {
	fmt.Println(google.Scheme)
}
```

## Status: Stable

This module is stable. It supports the [two most recent major
releases][go-support-policy] of Go.

Within those parameters, `must` follows semantic versioning. No
breaking changes will be made without incrementing the major version.

## Legal

Offered under the [MIT license][license].

[cmp]: https://pkg.go.dev/github.com/google/go-cmp/cmp
[go-support-policy]: https://golang.org/doc/devel/release#policy
[license]: https://github.com/akshayjshah/must/blob/main/LICENSE
[template.Must]: https://pkg.go.dev/text/template#Must
[regexp.MustCompile]: https://pkg.go.dev/regexp#MustCompile
