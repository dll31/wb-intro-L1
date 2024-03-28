// Задание 1: Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	age uint
}

func (h *Human) eat() {
	fmt.Println("Human is eating!")
}

func (h *Human) displayAge() {
	fmt.Printf("Your age is %v\n", h.age)
}

func (h *Human) sleep() {
	fmt.Println("Do not disturb! Human is sleeping!")
}

func (h *Human) DoSmth() {
	fmt.Println("Human is doing smth")
}

type Action struct {
	Human // встраиваем поля и методы Human в Action
}

// Переопределение метода Human.DoSmth()
func (a *Action) DoSmth() {
	fmt.Println("Wake up, Neo! Do something or follow the white rabbit!")
}

func (a *Action) coding() {
	fmt.Println("Coding!")
}

func (a *Action) calculateAge() {
	fmt.Printf("Trying to calculate your age. That`s so difficult! Ohh, your age %v!\n", a.age)
}

var (
	h Human
	a Action
)

func main() {
	// Доступ к методу Human.sleep()
	a.sleep()
	//out: Do not disturb! Human is sleeping!
	a.DoSmth()
	//out: Wake up, Neo! Do something or follow the white rabbit!
	h.DoSmth()
	// out: Human is doing smth

	// Получаем доступ к переменной Human.age
	a.age = 25
	// Проверим возраст через Human.displayAge()
	a.displayAge()
	//out: Your age is 25

	// Еще один способ получить доступ у переменной Human.age
	a.Human.age = 33
	// Проверим возраст через Action.calculateAge()
	a.calculateAge()
	// out: Trying to calculate your age. That`s so difficult! Ohh, your age 33!
}
