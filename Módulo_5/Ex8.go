/*
https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad

Description:
Given a set of numbers, return the additive inverse
of each. Each positive becomes negatives, and the
negatives become positives.

[1, 2, 3, 4, 5] --> [-1, -2, -3, -4, -5]
[1, -2, 3, -4, 5] --> [-1, 2, -3, 4, -5]
[] --> []

You can assume that all values are integers. Do not
mutate the input array.
*/

package main

import (
	"fmt"
)

func main() {

	rep1 := []int{1, 2, 3, 4, 5}
	rep2 := []int{1, -2, 3, -4, 5}
	rep3 := []int{0}
	fmt.Println(Invert(rep1))
	fmt.Println(Invert(rep2))
	fmt.Println(Invert(rep3))

}

func Invert(arr []int) []int {
	for i := range arr {
		arr[i] = -arr[i]
	}
	return (arr)
}
