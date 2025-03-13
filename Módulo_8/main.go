package main

import (
	"calculator/math" //nome_módulo/nome_pacote
	"fmt"

	"github.com/labstack/echo/v4" //porque usamos uma função de uma biblioteca externa
	// para externo, usar o go mod tidy e ele vai fazer o download
	//também cria automaticamente o arquivo go.sum, NUNCA ALTERAR ELE
)

func main() {

	// Echo instance
	e := echo.New()

	sum := math.Sum(1, 2)
	fmt.Println(sum)

	sub := math.sub(1, 2) //sub não vai funcionar pq a função começa com letra minúscula, ou seja, é privada
	fmt.Println(sub)
}
