# wildcard

[![Maintainability](https://api.codeclimate.com/v1/badges/50abc2cef12094116686/maintainability)](https://codeclimate.com/github/iamolegga/wildcard/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/50abc2cef12094116686/test_coverage)](https://codeclimate.com/github/iamolegga/wildcard/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/iamolegga/wildcard)](https://goreportcard.com/report/github.com/iamolegga/wildcard)

Golang's adaptation of [Java's Big-O(1) space solution](https://www.programcreek.com/2014/06/leetcode-wildcard-matching-java/).

Example:

```go
package main

import (
  "fmt"

  "github.com/iamolegga/wildcard"
)

func main() {
  fmt.Println(wildcard.Match("value", "v?l*"))
}
```
