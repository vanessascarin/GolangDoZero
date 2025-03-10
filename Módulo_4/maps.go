/*2-Maps: Heterogêneos
pode misturar tipos
estrutura chave-valor
[key] = value (a chave e a estrutura/valor têm um tipo específico)
chave tem um tipo, e o valor pode ter outro
map[k]v -> k=chave, v=valor

map[string]int
{"vanessa": 25, "bento": 4}
map[string]string
{"vanessa":"scarin,"bento":"cardoso"}*/

package main

import "fmt"

func main() {
	//inicializa map vazio
	idade := map[string]int{}
	idade["vanessa"] = 25
	idade["bento"] = 4
	fmt.Println("\n Printando Map completo:")
	fmt.Println(idade)
	fmt.Println("Printando Map pela chave:")
	fmt.Println(idade["vanessa"])
	fmt.Println(idade["bento"])

	fmt.Println("\n Agora, printando Map inicializado já com valores:")
	//inicializa map com valores
	anoNasc := map[string]int{
		"vanessa": 2000,
		"bento":   2001,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["vanessa"])
	fmt.Println(anoNasc["bento"])

	//adicionando novos elementos, como fazemos nas listas com append
	fmt.Println("\n Teste adicionando novos elementos:")
	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc)
}
