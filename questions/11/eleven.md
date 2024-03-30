Что выведет данная программа и почему?

```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```
Так как sync.WaitGroup нельзя передавать по значению, а только по указателю (компилятор предупреждает об этом при попытке передать wg по значению). Следовательно в каждую горутину попадает новый sync.WaitGroup и wg в main никогда не дождется освобождения т.о. программа упадет по deadlock, но значения от 0 до 5 скорее всего выведет.