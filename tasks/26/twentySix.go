/* Задача: Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

package main

import (
	"fmt"
	"unicode"
)

func allSymbolsUniqueCaseIndependent(str string) bool {
	letters := make(map[rune]struct{}, 0)
	var LCLetter rune

	for _, letter := range str {
		LCLetter = unicode.ToLower(letter)
		_, exist1 := letters[LCLetter]
		_, exist2 := letters[unicode.ToUpper(letter)]
		if exist1 || exist2 {
			return false
		} else {
			letters[LCLetter] = struct{}{}
		}
	}

	return true
}

func main() {

	str := []string{"abcd", "abCdefAaf", "aabcd"}

	for i := range str {
		fmt.Println(str[i], " - ", allSymbolsUniqueCaseIndependent(str[i]))
	}

}
