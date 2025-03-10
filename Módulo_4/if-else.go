package main

import "fmt"

func main() {
	numero := 1
	fmt.Println("\n Número =", numero)
	if numero == 1 { //true
		fmt.Println("Valor é igual a 1")
	} else {
		fmt.Println("O valor não é igual a 1")
	}

	numero = 7
	fmt.Println("\n Número =", numero)
	if numero == 1 { //true
		fmt.Println("Valor é igual a 1")
	} else if numero == 2 {
		fmt.Println("Valor é igual a 2")
	} else {
		fmt.Println("Valor é diferente de 1 e 2")
	}

	/*OPERAÇÕES
	  Soma: 1+1
	  Subtração: 2-1
	  Divisão: 10/2
	  Multiplicação: 2*2
	  Resto da divisão por x: 7%2 (resto da divisão por 2)*/

	fmt.Println("\n ##OPERAÇÕES")
	fmt.Println(1 + 1)
	fmt.Println(2 - 1)
	fmt.Println(2 * 2)
	fmt.Println(10 / 2)
	fmt.Println(10 % 2)

	numero = 8
	fmt.Println("\n Número =", numero)
	if numero%2 == 0 {
		fmt.Println(numero, " é par")
	} else {
		fmt.Println(numero, " é ímpar")
	}

	fmt.Println("\n \n Testando condicional usando strings e funções:")
	nome := "Vanessa"
	fmt.Println("\n Nome =", nome)
	acerta_nome(nome)

	nome = "Aline"
	fmt.Println("\n Nome =", nome)
	acerta_nome(nome)

}

func acerta_nome(primeironome string) {
	if primeironome == "Vanessa" {
		fmt.Println("Acertou meu nome! =D")
		return
	} else {
		fmt.Println("Este não é meu nome =(")
		return
	}
}
