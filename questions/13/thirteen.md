Что выведет данная программа и почему?

```go
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```
[100, 2, 3, 4, 5]. 

```go
type sliceHeader struct {
    Length        int
    Capacity      int
    ZerothElement *byte
}
``` 
Примерный header слайса. При передаче в функцию SomeAction происходит копирование данной структуры, но 
указатель ZerothElement все-равно указывает на слайс a из main, поэтому нулевой элемент изменяется.
Но при append происходит реалокация массива в памяти, так как len=capacity, и внутри локальной переменной 
v меняется ссылка на слайс. Следовательно слайс a видит только изменение нулевого элемента.  

Изменить можно вот так:

```go
func someAction(v *[]int8, b int8) {
	(*v)[0] = 100
	*v = append((*v), b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(&a, 6)
	fmt.Println(a)
}
```