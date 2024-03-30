Что такое интерфейсы, как они применяются в Go?

https://habr.com/ru/companies/vk/articles/463063/

* Интерфейсы в структурах:

    ```go
    type Printer interface {
        Print()
    }
    ```

    В данном примере мы определили интерфейс, с методом Print, теперь любая структура, 
    реализующая данный метод будет подходить под интерфейс Printer, например:

    ```go
    type SomeStructure struct {
        field1 string
        field2 string
    }

    func (s *SomeStructure) Print() {
        fmt.Printf("field1: %s\n field2: %s\n", s.field1, s.field2)
    }

    ```

    Если в какой-то метод будет передаваться тип Printer, то в качестве аргумента для 
    этой функции мы сможем передать SomeStructure, так как она реализует метод Print()

    ```go
    func printFields(p *Printer) {
        p.Print()
    }

    structure := SomeStruct{
        field1: "first field",
        field2: "second field",
    }

    printFields(&structure)

    ```

* interface{} в переменной

    Интрефейс в переменной удовлетворяет объекту любого типа, например:
    ```go
    person := make(map[string]interface{}, 0)

    person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64
    ```

    Но при попытке сделать
    ```go
    person["age"] = person["age"] + 1
    ```
    Увидим ошибку
    ```
    invalid operation: person["age"] + 1 (mismatched types interface {} and int)
    ```
    так как значения, хранящиеся в map потеряли свой исходный тип и преобрели тип interface{}, для работы с person["age"], как с int, необходимо привести его к типу int.
    ```go
    age, ok := person["age"].(int)
    if !ok {
        log.Fatal("could not assert value to int")
        return
    }

    person["age"] = age + 1
    ```

    В данном виде интерфейс состоит из двух полей - (value, type). Про определение типов в runtime из interface{} было в задаче 14.
