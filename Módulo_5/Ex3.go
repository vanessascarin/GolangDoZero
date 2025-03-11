/*
https://www.codewars.com/kata/5715eaedb436cf5606000381

Task
You get an array of numbers, return the sum of all of
 the positives ones.

Example
[1, -4, 7, 12] => 1+7+12=20
Note
If there is nothing to sum, the sum is default to 0.
*/

package main

import "fmt"

func main() {
	var soma int
	numbers := []int{1, -4, 7, 12}
	soma = PositiveSum(numbers)
	fmt.Println(soma)

}

func PositiveSum(numbers []int) int {

	var soma int
	tamanho := len(numbers)
	for i := 0; i < tamanho; i++ {
		if numbers[i] > 0 {
			soma += numbers[i]
		}
	}
	return soma

}
