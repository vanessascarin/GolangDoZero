/*
https://www.codewars.com/kata/580755730b5a77650500010c

Description:
Given a string s, your task is to return another
string such that even-indexed and odd-indexed
characters of s are grouped and the groups are
space-separated. Even-indexed group comes as first,
 followed by a space, and then by the odd-indexed
 part.

Examples
input:    "CodeWars" => "CdWr oeas"
           ||||||||      |||| ||||
indices:   01234567      0246 1357

Even indices 0, 2, 4, 6, so we have "CdWr" as the
 first group.
Odd indices are 1, 3, 5, 7, so the second group is
 "oeas".
And the final string to return is "Cdwr oeas".

Notes
Tested strings are at least 8 characters long.

*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	teste1 := "CodeWars"
	teste2 := "YCOLUE'VREER"

	fmt.Println(SortMyString(teste1))
	fmt.Println(SortMyString(teste2))
}

func SortMyString(s string) string {
	LetrasLista := strings.Split(s, "")
	var OddIndex, EvenIndex, resultado string
	for i := 0; i < len(LetrasLista); i++ {
		if i%2 == 0 {
			EvenIndex += LetrasLista[i]
		} else {
			OddIndex += LetrasLista[i]
		}
		resultado = fmt.Sprintf("%s %s", EvenIndex, OddIndex)
	}

	return resultado
}
