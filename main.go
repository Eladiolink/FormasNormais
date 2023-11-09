package main

import (
	formagreibach "FormasNormais/formasNormais/formaGreibach"
	"FormasNormais/helpers"
	"FormasNormais/helpers/export"
	"FormasNormais/helpers/gramatica"
	"fmt"
	"os"
)

func main() {
	/* A funçãp helpers.File pega a localização do arquivo no segundo argumento do cli,
	e converte para uma estrutura gramatica e retorna o ponteiro dela */
	gramatica := helpers.File()

	// simplificacao.Simplificacao(gramatica)
	// formaprechomsky.Formaprechomsky(gramatica)

	gramatica = controller(gramatica)
	export.ExportToJson(gramatica)
}

func controller(gramatica *gramatica.Gramatica) *gramatica.Gramatica {
	switch os.Args[2] {
	case "-C":
		return gramatica
	case "-G":
		return formagreibach.FormaGreibach(gramatica)
	default:
		fmt.Println("Argunmento de entrada inválido!")
	}

	return gramatica
}
