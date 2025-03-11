package main

import (
	"fmt"
)

/* loops
laços de repetição
repetir tarefas*/

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
		/* essa é a mesma coisa que: sum = sum +1
		sum -=i -> sum = sum - 1*/
		//toda variável criada dentro do for só existe dentro do escopo do for
	}
	fmt.Println(sum)

	//loop infinito
	/*for {
		fmt.Println("Golang do zero") //usar ctrl + c para parar de executar o código infinito
		time.Sleep(2 * time.Second) //congela o código na tela por 2 segundos
	}*/

	//for range - para percorrer uma lista
	frutas := []string{"laranja", "maça", "banana", "uva", "kiwi"}
	for _, fruta := range frutas { //aqui demos o nome de "fruta" para cada elemento que será percorrido na lista, e o _ é a posição
		fmt.Println(fruta)
	}

	turma := []string{"monica", "magali", "cascão", "cebolinha"}
	for i, crianca := range turma { //também podemos colocar o i no lugar do _ caso o valor do índice for ser utilizado
		fmt.Println("Criança: ", crianca, "Índice: ", i)
	}
}
