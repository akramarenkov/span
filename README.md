# Span

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/span.svg)](https://pkg.go.dev/github.com/akramarenkov/span)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/span)](https://goreportcard.com/report/github.com/akramarenkov/span)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/span/badge.svg)](https://coveralls.io/github/akramarenkov/span)

## Purpose

Library that allows you to divide a sequence of something into spans

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/span"
)

func main() {
    spans, err := span.Int(1, 8, 3)
    fmt.Println(err)
    fmt.Println(spans)
    // Output:
    // <nil>
    // [{1 3} {4 6} {7 8}]
}
```
