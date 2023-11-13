package simplificacao

import (
	"FormasNormais/helpers"
	"FormasNormais/helpers/gramatica"
	"fmt"
)

func Simplificacao(gramatica *gramatica.Gramatica) {
	gramatica.P = RemocaoPalavraVazia(gramatica)
	RemocaoUnitarias(gramatica)
	helpers.PrintProducoes(gramatica)
	RemocaoInuteis(gramatica)

	fmt.Println("\nGRAMÁTICA SIMPLIFICADA!!! （っ＾▿＾）っ")
	fmt.Println()
}
