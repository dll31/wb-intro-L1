Что выведет данная программа и почему?

```go
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```
Выведется 0, так как в if инициализируется локальная переменная с именем n и внутри if инкрементация произойдет, 
но глобальная переменная n останется равной 0.