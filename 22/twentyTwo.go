// Задание: Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	f := new(big.Int)
	f.SetString("123402349309203090293090492090904902930903942390920934", 10)

	s := new(big.Int)
	s.SetString("32932287929817237975297172935987237498", 10)

	t := new(big.Int)
	t.Add(f, s)

	fmt.Println("x + y = ", t)

	t.Sub(f, s)
	fmt.Println("x + y = ", t)

	t.Mul(f, s)
	fmt.Println("x * y = ", t)

	t.Div(f, s)
	fmt.Println("x / y = ", t)

}
