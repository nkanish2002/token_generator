# Token Generator

Token Generator is a Golang based unique string generator which makes 
use of the "math/rand" package.

## Installation

```sh
go get github.com/nkanish2002/token_generator
```

## Usage

```go
package main

import (
	"github.com/nkanish2002/token_generator"
	"fmt"
)

func main() {
	g := token_generator.Generator{}
	g.New()
	for i := 0; i < 100; i++ {
			token := g.GetToken(32)
			fmt.Println(token)
	}
}
```

## Note

This is my first go library, I do not guarentee any performance benchmarks or bugs.

For more complete documentation follow https://godoc.org/github.com/nkanish2002/token_generator
