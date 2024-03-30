/* Задача: К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

/*
Проблема в том, что justString ссылается на v, то есть вся строка v хранится в памяти и не может быть удалена.
Если сделать copy(justString, v[:100]), то нельзя гаранировать, что первые 100 символов войдут в срез, так как каждый символ может кодироваться 1 - 4 байтами
*/

package main

import (
	"fmt"
	"math/rand"
)

var justString string

func createHugeString(lenght int) string {
	const letters = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNMйцукенгшщзхъфывапролджэячсмитьбюЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮ"
	constRunes := []rune(letters)

	out := make([]rune, lenght)
	for i := 0; i < lenght; i++ {
		out[i] = constRunes[rand.Intn(len(constRunes))]
	}

	return string(out)

}

func someFunc() {
	v := createHugeString(1024) // будет удалена после выхода из функции

	out := make([]rune, 100)

	// Производим поэлементное копирование ста первых символов
	i := 0
	for _, letter := range v {
		out[i] = letter
		i++
		if i == 100 {
			break
		}
	}

	justString = string(out)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
