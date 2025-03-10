package main

import "fmt"

func main() {
	fmt.Println(soma(42, 13))

	nome := printaNome("VANESSA")
	fmt.Println(nome)

	nome1, nome2 := dois_retornos("HELENA")
	fmt.Println(nome1)
	fmt.Println(nome2)

	//caso apenas um dos retornos seja utilizado, como "burlar" os outros:
	fruta1, fruta2 := burlar_segundo_retorno("BANANA", "LARANJA")
	//se a linha acima fosse:
	//fruta1 := burlar_segundo_retorno("BANANA", "LARANJA")
	//daria erro de assignment mismatch, pois retorna 2 valores e só está atribuindo à uma variável da main
	//pra mostrar como resolver isso primeiro printaremos os dois valores, somente para demonstrar:
	fmt.Println(fruta1)
	fmt.Println(fruta2)
	//agora sim burlando para atribuir somente à uma variável da main, usamos o UNDERLINE (_):
	fruta1, _ = burlar_segundo_retorno("MORANGO", "ABACAXI")
	fmt.Println(fruta1)

}

func soma(x int, y int) int {
	return x + y
}

func printaNome(nome string) string {
	return nome
}

//função que retorna 2 valores:

func dois_retornos(nome string) (string, string) {
	return nome, nome
}

//funcao para utilizar somente um dos retornos
func burlar_segundo_retorno(comida1 string, comida2 string) (string, string) { /*também poderia ser (comida1, comida2 string)*/
	return comida1, comida2
}

/*Função começando com letra minúscula: PRIVADA
Só pode ser utilizada no próprio pacote
PÚBLICA: pode ser usada fora do pacote, por exemplo a função Println, do pacote fmt, que usamos assim: fmt.Println*/
