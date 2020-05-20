# wildcard

[![GoDoc](https://godoc.org/github.com/iamolegga/wildcard?status.svg)](https://godoc.org/github.com/iamolegga/wildcard)
[![Maintainability](https://api.codeclimate.com/v1/badges/50abc2cef12094116686/maintainability)](https://codeclimate.com/github/iamolegga/wildcard/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/50abc2cef12094116686/test_coverage)](https://codeclimate.com/github/iamolegga/wildcard/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/iamolegga/wildcard)](https://goreportcard.com/report/github.com/iamolegga/wildcard)

Package `wildcard` provides utils for working with a classic wildcard matching (only `?` and `*` are supported).

`Match` function is a golang's adaptation of [Java's O(1) space solution](https://www.programcreek.com/2014/06/leetcode-wildcard-matching-java/) and ~ 3.37 times faster (76.0ns/22.5ns) then standard `filepath.Match`.

## Example

```go
package main

import (
  "fmt"

  "github.com/iamolegga/wildcard"
)

func main() {
  fmt.Println(wildcard.Match("value", "v?l*"))
  fmt.Println(wildcard.IsPattern("v?l*"))
}
```

## Performance

It was compared with `filepath.Match` function with several runs:

```sh
$ for i in {1..10}; do go test -run=NONE -bench=. ./... >> old.txt; done
$ # change implementation to `filepath.Match`
$ for i in {1..10}; do go test -run=NONE -bench=. ./... >> new.txt; done
$
$ benchstat old.txt new.txt
name     old time/op  new time/op  delta
Match-4  22.5ns ± 2%  76.0ns ±19%  +238.01%  (p=0.000 n=10+10)
```

`old` - current implementation

`new` - utilize `filepath.Match`
