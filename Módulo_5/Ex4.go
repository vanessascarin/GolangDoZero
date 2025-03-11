/*
https://www.codewars.com/kata/576bb71bbbcf0951d5000044

Given an array of integers.

Return an array, where the first element is the count
 of positives numbers and the second element is sum
 of negative numbers. 0 is neither positive nor
 negative.

If the input is an empty array or is null, return an
 empty array.

Example
For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12,
 -13, -14, -15], you should return [10, -65].
*/

package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	resposta := CountPositivesSumNegatives(numbers)
	fmt.Println(resposta)

}

func CountPositivesSumNegatives(numbers []int) []int {
	var count, sum int
	for _, number := range numbers {
		if number > 0 {
			count++
		} else {
			sum += number
		}
	}
	return []int{count, sum} // your code here
}
