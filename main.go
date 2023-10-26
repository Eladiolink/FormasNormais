package main

import (
	"FormasNormais/helpers"
	"FormasNormais/helpers/simplificacao"
	"fmt"
	"os"
)

func main() {
	// A funçãp helpers.File pega a localização do arquivo no segundo argumento do cli, e converte para uma estrutura gramatica e retorna o ponteiro dela
	gramatica := helpers.FileTxt()

	fmt.Println(gramatica.Regras)
	simplificacao.Simplificacao(gramatica)
}

func controller() {
	switch os.Args[2] {
	case "C":
		fmt.Println("Forma Normal Chomsky")
	case "G":
		fmt.Println("Forma Normal Geibach")
	default:
		fmt.Println("Argunmento de entrada inválido!")
	}
}
