Что выведет данная программа и почему?

```go
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```

Значение p не изменяется, потому что в функцию update передается копия p. 
Если нужно изменить значение под указателем p, то можно сделать следующее:
```go
// функция принимает указатель на переменную p
func update(p **int) {
	b := 2
	*p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(&p)
	fmt.Println(*p)
}
```