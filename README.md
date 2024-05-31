# Zero

[![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](LICENSE)

**Zero** is a very simple package that takes any input of any data type and zeros it from memory, making it unusable and unreadable in memory.

This is especially useful for when you're dealing with sensitive data that needs to be quickly removed from memory.

This package also provides plenty of tests to ensure that each function works as expected.

## Example

```go
package main

import (
	"fmt"

	"github.com/notwithering/zero"
)

func main() {
	// data you do not want to be stored into memory for long
	var info string = "my super secret password"

	// zero out the data
	zero.Zero(&info)

	// there is no data to even print
	fmt.Println(info)
}
```
