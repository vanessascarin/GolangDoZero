package main

import "fmt"

/*Structs
Forma de criar sua própria estrutura de dados
Personalizar de acordo com a sua necessidade
Podemos usar vários tipos diferentes

type <nome da nossa estrutura> struct { <campos> }*/
type Pessoa struct {
	Nome  string
	Idade int
}

func main() {

	fmt.Println("\n Structs sendo exibidas e valores sendo atribuídos aos campos simultaneamente:")
	fmt.Println(Pessoa{"Vanessa", 25})
	fmt.Println(Pessoa{Nome: "Bento", Idade: 4}) //Forma mais aconselhável de escrever a atribuição
	fmt.Println(Pessoa{Nome: "Vanessa"})         //Também podemos criar uma estrutura sem todos os campos preenchidos (será 0)

	//Atrelando a struct a uma variável, que agr vira uma variável do tipo struct
	p1 := Pessoa{Nome: "Bob", Idade: 2}
	fmt.Println("\n Campos de structs que estão atribuídas a uma variável:")
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	//Alterando o valor de um campo
	p1.Idade = 3
	fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Patrick", Idade: 2}

	//Criando um slice de pessoas:
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println("\n Structs que fazem parte de um slice:")
	fmt.Println(pessoas)

	//structs + map
	alunos := map[string][]Pessoa{}
	alunos["programação"] = pessoas
	fmt.Println("\n Map chamado cuja chave é "programação" e seu valor é o slice (a lista) composto por structs Pessoa:")
	fmt.Println(alunos)

	fmt.Println("\n Maps da mesma forma, mas com structs sendo exibidas e valores sendo atribuídos aos campos simultaneamente:")
	var estudantes = map[string][]Pessoa{
		"Estatística": {{Nome: "Maria", Idade: 68}, {Nome: "Joaquim", Idade: 72}},
		"Engenharia":  {{Nome: "Maria2", Idade: 68}, {Nome: "Joaquim2", Idade: 72}},
	}

	fmt.Println(estudantes)

	//struct herdando campos de outra struct

	type Profissao struct {
		Pessoa
		Tipo string
	} // todas as declarações de struct devem preferencialmente estar juntas no topo no código, mas por motivos didáticos esta cabe melhor aqui

	prof := Profissao{p2, "dev"}
	fmt.Println("\n Struct que herda campos de outra struct")
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)

	fmt.Println("\n")
}
