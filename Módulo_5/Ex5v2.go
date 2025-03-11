/*
\https://www.codewars.com/kata/576bb71bbbcf0951d5000044https://www.codewars.com/kata/544675c6f971f7399a000e79

We need a function that can transform a string into a
 number. What ways of achieving this do you know?

Note: Don't worry, all inputs will be strings, and
 every string is a perfectly valid representation of
  an integral number.

Examples
"1234" --> 1234
"605"  --> 605
"1405" --> 1405
"-7" --> -7
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	entrada_string1 := "1234"
	entrada_string2 := "605"
	entrada_string3 := "1405"
	entrada_string4 := "-7"
	fmt.Println(StringToNumber(entrada_string1))
	fmt.Println(StringToNumber(entrada_string2))
	fmt.Println(StringToNumber(entrada_string3))
	fmt.Println(StringToNumber(entrada_string4))
}

func StringToNumber(entrada string) int {
	saida_inteiro, _ := strconv.Atoi(entrada) // burlando o erro com _
	return (saida_inteiro)
}
