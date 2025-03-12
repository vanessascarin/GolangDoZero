/*
/*
https://www.codewars.com/kata/53da3dbb4a5168369a0000fe

I have a cat and a dog.
I got them at the same time as kitten/puppy.
That was [humanYears] years ago.
Return their respective ages now as
[humanYears,catYears,dogYears]

NOTES:

humanYears >= 1
humanYears are whole numbers only

Cat Years
15 cat years for first year
+9 cat years for second year
+4 cat years for each year after that

Dog Years
15 dog years for first year
+9 dog years for second year
+5 dog years for each year after that
*/

package main

import (
	"fmt"
)

func main() {

	teste1 := 1
	teste2 := 2
	teste3 := 10

	fmt.Println(CalculateYears(teste1))
	fmt.Println(CalculateYears(teste2))
	fmt.Println(CalculateYears(teste3))

}

/*func CalculateYears(years int) (result [3]int) {
	result[0] = years
	result[1], result[2] = 15, 15
	if years > 1 {
		result[1] += 9
		result[2] += 9
	}
	for i := 3; i <= years; i++ {
		result[1] += 4
		result[2] += 5
	}
	return (result)
}*/

//ou não colocar esse segundo for, e sim:

func CalculateYears(years int) (result [3]int) {
	result[0] = years
	result[1], result[2] = 15, 15
	if years > 1 {
		result[1] += 9 + (years-2)*4
		result[2] += 9 + (years-2)*5
	}
	return /*não precisa escrever o (result) aqui
	porque o nome da variável de retorno já está
	escrito na primeira linha da função*/
}
