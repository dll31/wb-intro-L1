Какой размер у структуры struct{}{}?

0 байт. Для проверки можно запустить следующий код
```go
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    s := struct{}{}

    fmt.Printf("s size: %v\n", unsafe.Sizeof(s)) //s size: 0
}
```
