package main

import (
	formachomsky "FormasNormais/formasNormais/formaChomsky"
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

	gramatica = controller(gramatica)
	export.ExportToJson(gramatica)
}

func controller(gramatica *gramatica.Gramatica) *gramatica.Gramatica {
	switch os.Args[2] {
	case "-C":
		return formachomsky.FormaChomsky(gramatica)
	case "-G":
		return formagreibach.FormaGreibach(gramatica)
	default:
		fmt.Println("Argunmento de entrada inválido!")
	}

	return gramatica
}
