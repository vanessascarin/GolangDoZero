/*LISTAS

1-Arrays e Slices: Homogêneos
todos os elementos têm o mesmo tipo
[1, 2, 3, 4, 5, 6] - []int
["vanessa", "bento", "golang"] - []strings
[1.001, 2.23, 3.9, 4.322, 5, 6] - []float64

2-Maps: Heterogêneos
pode misturar tipos
estrutura chave-valor
[key] = value (a chave e a estrutura/valor têm um tipo específico)
chave tem um tipo, e o valor pode ter outro
map[string]int
{"vanessa": 25, "bento": 4}
map[string]string
{"vanessa":"scarin,"bento":"cardoso"}

ARRAY
- Tamanho fco, de zero ou mais elementos do mesmo tipo
- Acessamos os valores com índice: a[0], a[1]
- Função embutida len() retorna o tamanho do array
- Por conta do tamanho fixo, não é tão usado, só em casos específicos

SLICE
- Tipo o array, mas sem tamanho fixo
- Acessamos os valores com índice: a[0], a[1]
- Função embutida len() retorna o tamanho do array
- Função append() usada para adicionar valores no slice*/

package main

import "fmt"

func main() {
	//Array - tamanho fixo
	fmt.Println("### ARRAY ###")
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println("* Chamando por índice:")
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[0], array[1])
	fmt.Println("* Chamando pelo array todo:")
	fmt.Println(array)

	fmt.Println("* Exemplo de array declarado já com valores -")
	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("* Chamando pelo array todo:")
	fmt.Println(numPrimos)
	fmt.Println("* Chamando por índice:")
	fmt.Println(numPrimos[0])
	fmt.Println(numPrimos[1])
	fmt.Println("* Chamando por segmento:")
	fmt.Println(numPrimos[0:3]) //mostra valores da posição 0 até a 3
	fmt.Println(numPrimos[:1])  // mostra valores antes do índice 1
	fmt.Println(numPrimos[1:])  // mostra valores depois do índice 1
	fmt.Println("\n\n")

	//Slice
	fmt.Println("### SLICE ###")
	//var slice []string -> DÁ ERRO, usar make
	slice := make([]string, 5) //coloca um tamanho inicial mas ele não é fixo, aceitará valores a mais
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println("* Chamando por índice:")
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	slice[2] = "Vanessa"
	fmt.Println(slice[2])

	fmt.Println("* Tamanho do slice:")
	fmt.Println(len(slice)) //printa o tamanho do slice
	fmt.Println("* Posições vazias do slice:")
	fmt.Println(slice[3])
	fmt.Println(slice[4]) //posições que estão vazias

	//se for criar o slice já com valores, não precisa do make

	fmt.Println("* Exemplo de slice declarado já com valores -")
	numPares := []int{2, 3, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14)
	fmt.Println(numPares)

}
