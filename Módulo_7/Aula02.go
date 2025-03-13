package main

import "fmt"

type Cachorro struct {
	Raca  string
	Cor   string
	Patas int
}

type Gato struct {
	Cor   string
	Patas int
}

func (c Cachorro) Barulho() string {
	return "au au"
}

func (c Cachorro) NumeroDePatas() int {
	return c.Patas
}

func (g Gato) Barulho() string {
	return "miau"
}

func (g Gato) NumeroDePatas() int {
	return g.Patas
}

//interface é o conjunto de métodos que identificam se a struct é do tipo agrupado

type Animal interface { //ou seja, pra struct Gato fazer parte da interface Animal, ela precisa obedecer, conter, esses métodos atrelados à struct
	Barulho() string //lembrar de definir que tipo esse método recebe, nome é opcional
	NumeroDePatas() int
}

func FazBarulho(animal Animal) {
	fmt.Println(animal.Barulho())
}

type Pessoa struct {
	Nome             string
	Idade            int
	Prof             string
	TempoDeProfissao int
}

type Crianca struct {
	Nome  string
	Idade int
}

func (c Crianca) FalaBomDia() string {
	return c.Nome + "deseja bom dia pra você!"
}

func (p Pessoa) FalaBomDia() string {
	return fmt.Sprintf("%s deseja bom dia pra você!", p.Nome)
}

func (p Pessoa) Profissao() string {
	return fmt.Sprintf("%s tem %d anos como %s", p.Nome, p.TempoDeProfissao, p.Prof)
}

type Adulto interface {
	FalaBomDia() string
	Profissao() string
}

func BomDia(adulto Adulto) string {
	return adulto.FalaBomDia()
}

func main() {

	adulto := Pessoa{
		Nome:             "vanessa",
		Idade:            25,
		Prof:             "dev",
		TempoDeProfissao: 1,
	}

	crianca := Crianca{
		Nome:  "magali",
		Idade: 10,
	}

	crianca.FalaBomDia()
	adulto.FalaBomDia()
	//BomDia(crianca) --> vai dar erro pq esse é um método da interface adulto, e criança não faz parte dela pq não cumpre o método de profissão.
	BomDia(adulto)

	cachorro := Cachorro{
		Raca:  "vira lata",
		Cor:   "caramelo",
		Patas: 4,
	}

	gato := Gato{
		Cor:   "branco",
		Patas: 4,
	}

	fmt.Println(cachorro.Barulho())
	fmt.Println(gato.Barulho())
	fmt.Println(cachorro.NumeroDePatas())
	fmt.Println(gato.NumeroDePatas())

	fmt.Println("Agora printando direto pelas interfaces sem chamar as 2 funções:")
	FazBarulho(gato)
	FazBarulho(cachorro)
}
