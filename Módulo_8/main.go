package main

import (
	"calculator/math" //nome_módulo/nome_pacote
	"fmt"
)

func main() {
	sum := math.Sum(1, 2)
	fmt.Println(sum)

	sub := math.Sub(1, 2)
	fmt.Println(sub)
}
