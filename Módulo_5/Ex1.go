/*
	https://www.codewars.com/kata/55685cd7ad70877c23000102

In this simple assignment you are given a number and

	have to make it negative. But maybe the number is
	already negative?
*/
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Digite um número: ")
	fmt.Scan(&num)
	fmt.Println("Você digitou:", num)
	if num > 0 {
		num = make_negative(num)
		fmt.Println("O negativo deste número é:", num)
	} else if num == 0 {
		fmt.Println("O número é zero.")
	} else {
		fmt.Println("O número já é negativo.")
	}
}

func make_negative(num_original int) int {
	negativo := (num_original * (-1))
	return negativo
}
