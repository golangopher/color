# How to use

```go
package main

import (
	"fmt"
	"github.com/golangopher/color"
)

func main() {
	fmt.Println(color.C("blue", "I'm blue"))
	fmt.Printf("This has no color but %s", color.C("green", "this is green"))
	// ...
}
```

# Colors supported

black, red, green, yellow, blue, cyan, white, magenta(default)

# OS

- Linux: yes
- Windows: powershell
- Others: ?