Какой самый эффективный способ конкатенации строк?

через strings.Builder

```go

package main

import (
	"fmt"
	"strings"
)

func main() {

    someStrings := []string{"first string", "second string", "third string"}

    var builder strings.Builder

    // Если известен конечный размер строки. Так будет работать быстрее
    //builder.Grow(len(str))

    for i := range someStrings {
        builder.WriteString(someStrings[i])    
    }

    fmt.Println(builder.String())
}

```
