package main

import "fmt"

// Ponteiros: Um ponteiro nada mais é do que uma variável que ao invés
// de armazenar um valor (1, "olá", true..), armazena um end de memória

func zeroValue(i int) {
	i = 0
	fmt.Println("End de memório do i dentro da func:", &i)
}

func zeroPointer(i *int) {
	*i = 0
}

func main() {
	i := 1
	fmt.Println("Valor inicial:", i)
	fmt.Println("Valor end de memória:", &i)

	zeroValue(i)
	fmt.Println("zeroval:", i)

	zeroPointer(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}
