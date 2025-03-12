/* NOVO NÍVEL DE DIFICULDADE
https://www.codewars.com/kata/56f3a1e899b386da78000732/train/go

Write a function partlist that gives all the ways to
 divide a list (an array) of at least two elements
 into two non-empty parts.

- Each two non empty parts will be in a pair (or an
 array for languages without tuples or a structin
 C - C: see Examples test Cases - )
- Each part will be in a string
- Elements of a pair must be in the same order as
in the original array.

*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	teste1 := []string{"I", "wish", "I", "hadn't", "come"}
	teste2 := []string{"cdIw", "tzIy", "xDu", "rThG"}

	fmt.Println(PartList(teste1))
	fmt.Println(PartList(teste2))
}

func PartList(arr []string) string {
	tam := len(arr)
	var divisao, possibilidades string
	for i := 1; i < tam; i++ {
		parte1 := strings.Join(arr[:i], " ")
		parte2 := strings.Join(arr[i:], " ")
		divisao = parte1 + ", " + parte2               //nessa e na linha debaixo são formas diferentes de fazer a mesma coisa:
		possibilidades += fmt.Sprintf("(%s)", divisao) //formatar a entrada de strings (com espaço, vírgula, parênteses, etc)
	}
	return (possibilidades)
}
