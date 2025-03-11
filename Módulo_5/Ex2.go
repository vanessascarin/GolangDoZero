/*
https://www.codewars.com/kata/55d24f55d7dd296eb9000030

Write a program that finds the summation of every
number from 1 to num. The number will always be a
positive integer greater than 0. Your function only
needs to return the result, what is shown between
parentheses in the example below is how you reach
that result and it's not part of it, see the sample
 tests.

For example (Input -> Output):

2 -> 3 (1 + 2)
8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)
*/

package main

import "fmt"

func main() {
	var num, soma int
	fmt.Println("Digite um número inteiro maior que zero: ")
	fmt.Scan(&num)
	fmt.Println("Você digitou:", num)

	if num <= 0 {
		fmt.Println("Por favor, não digite números negativos ou zero. Insira o número novamente:")
		fmt.Scan(&num)
	}

	for i := 1; i <= num; i++ {
		soma += i
	}

	fmt.Println("A soma dos números de 1 até", num, "é:", soma)

}
