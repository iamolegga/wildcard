# wildcard

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
