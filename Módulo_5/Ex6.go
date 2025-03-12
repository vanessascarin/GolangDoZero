/*
https://www.codewars.com/kata/5265326f5fda8eb1160004c8

Description:
We need a function that can transform a number
(integer) into a string.

What ways of achieving this do you know?

Examples (input --> output):
123  --> "123"
999  --> "999"
-100 --> "-100"
*/

package main

import (
	"fmt"
)

func main() {

	entrada_int1 := 67
	entrada_int2 := 79585
	entrada_int3 := 3
	entrada_int4 := -1
	entrada_int5 := 0
	fmt.Println(NumberToString(entrada_int1))
	fmt.Println(NumberToString(entrada_int2))
	fmt.Println(NumberToString(entrada_int3))
	fmt.Println(NumberToString(entrada_int4))
	fmt.Println(NumberToString(entrada_int5))

}

/*
func NumberToString(n int) string {
	return (strconv.Itoa(n)) //como esse comando só tem 1 retorno, podemos colocá-la direto no return
}

ALTERNATIVA:*/

func NumberToString(n int) string {
	return (fmt.Sprintf("%d", n)) //formata qualquer coisa pra string
}
