# Token Generator

Token Generator is a Golang based unique string generator which makes 
use of the "math/rand" package.

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


