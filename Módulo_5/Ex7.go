/*
https://www.codewars.com/kata/57a0e5c372292dd76d000d7e

Description:
Write a function that accepts a non-negative
integer n and a string s as parameters, and returns
 a string of s repeated exactly n times.

Examples (input -> output)
6, "I"     -> "IIIIII"
5, "Hello" -> "HelloHelloHelloHelloHello"
*/

package main

import (
	"fmt"
	"strings" //pra alternativa
)

func main() {

	rep1 := 4
	rep2 := 3
	rep3 := 2
	rep4 := 0
	rep5 := 0
	rep6 := 5
	rep7 := 6
	rep8 := 5
	palavra1 := "a"
	palavra2 := "hello"
	palavra3 := "abc"
	palavra4 := " "
	palavra5 := "I"
	palavra6 := " "
	palavra7 := "I"
	palavra8 := "Hello"
	fmt.Println(RepeatStr(rep1, palavra1))
	fmt.Println(RepeatStr(rep2, palavra2))
	fmt.Println(RepeatStr(rep3, palavra3))
	fmt.Println(RepeatStr(rep4, palavra4))
	fmt.Println(RepeatStr(rep5, palavra5))
	fmt.Println(RepeatStr(rep6, palavra6))
	fmt.Println(RepeatStr(rep7, palavra7))
	fmt.Println(RepeatStr(rep8, palavra8))

}

/*func RepeatStr(repetitions int, value string) string {
	var resultado string
	for i := 0; i < repetitions; i++ {
		resultado += value
	}
	return (resultado)
}

ALTERNATIVA:*/

func RepeatStr(repetitions int, value string) string {
	return (strings.Repeat(value, repetitions)) //função pronta
}
