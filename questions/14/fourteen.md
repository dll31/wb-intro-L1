Что выведет данная программа и почему?

```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```

[b b a][a a]
После append слайс внутри анонимной функции ссылается не на исходный слайс, а на новый, 
поэтому дальнейшие его изменения не видны в исходном