package main

import "fmt"

// Ponteiros: Um ponteiro nada mais é do que uma variável que ao invés
// de armazenar um valor (1, "olá", true..), armazena um end de memória

func main() {
	// Memória -> essa memória tem um endereço -> esse endereço guarda um valor

	i := 1
	fmt.Println("Valor inicial:", i)
	fmt.Println("Valor end de memória:", &i) //&var -> apontando para o end de memória

	a := &i
	fmt.Println("Variavel a:", a)
	fmt.Println("Variavel *a:", *a) // quando uso *, estamos falando do endereço! E não do valor
	fmt.Println("Variavel &a:", &a)
}
