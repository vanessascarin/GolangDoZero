package main

import "fmt"

func main() {
	//Variaveis
	//var + nome da variável + tipo
	var nome string
	nome = "bento"
	fmt.Println(nome)

	nome = "vanessa"
	fmt.Println(nome)

	var idade int
	idade = 4
	fmt.Println(idade)

	var a = "vanessa"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	//Constantes
	const idadeBento = 4
	fmt.Println(idadeBento)
}
