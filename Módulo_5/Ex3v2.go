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
	for _, number := range numbers {
		if number > 0 {
			soma += number
		}
	}
	return soma

}
