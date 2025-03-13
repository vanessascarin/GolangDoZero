package main

import "fmt"

type Pessoa struct {
	Nome      string
	Idade     int
	Profissao string
}

/*func ListaNomeEIdade(nome string, idade int) string {
	return fmt.Sprintf("%s tem %d anos", nome, idade)
} ----> ou então: */

/*func ListaNomeEIdade(pessoa Pessoa) string { //melhor pq passa menos parâmetros
	return fmt.Sprintf("%s tem %d anos", pessoa.Nome, pessoa.Idade)
} ----> MAS COMO ASSOCIAR ESSE MÉTODO COM A STRUCT PESSOA? */

func (p Pessoa) ListaNomeEIdade() string { //usando o receiver, passando o identificar de só 1 ou 2 letras (no caso, p) e não passa os parâmetros mais
	return fmt.Sprintf("%s tem %d anos", p.Nome, p.Idade)
}

func main() {
	pessoa := Pessoa{
		Nome:      "Vanessa",
		Idade:     25,
		Profissao: "dev",
	}

	//println(ListaNomeEIdade(pessoa.Nome, pessoa.Idade)) -> ou então
	//println(ListaNomeEIdade(pessoa)) ---> E SE O MÉTODO ESTIVER ATRELADO A UMA STRUCT?
	println(pessoa.ListaNomeEIdade())
}
