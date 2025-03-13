package math //pode ter o mesmo nome que o arquivo ou não

func Sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

//usar o go mod init e um "identificador do módulo", por
// exemplo github.com/vanessascarin/calculator, se eu
// tivesse um projeto no github com esse path
// NESSE CASO, usamos go mod init calculator, e ele cria o arquivo
// go.mod no mesmo diretório que o código do pacote
