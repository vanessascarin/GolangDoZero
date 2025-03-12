/*
https://www.codewars.com/kata/53da3dbb4a5168369a0000fe

Description:
Create a function that takes an integer as an
argument and returns "Even" for even numbers or
"Odd" for odd numbers.
*/

package main

import (
	"fmt"
)

func main() {

	teste1 := 1
	teste2 := 2
	teste3 := -1
	teste4 := -2
	teste5 := 0
	fmt.Println(EvenOrOdd(teste1))
	fmt.Println(EvenOrOdd(teste2))
	fmt.Println(EvenOrOdd(teste3))
	fmt.Println(EvenOrOdd(teste4))
	fmt.Println(EvenOrOdd(teste5))

}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return ("Even")
	}
	return ("Odd")
}
