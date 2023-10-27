package simplificacao

import (
	"FormasNormais/helpers/gramatica"
	"fmt"
)

func Simplificacao(gramatica *gramatica.Gramatica) {
	gramatica.P = RemocaoPalavraVazia(gramatica)
	RemocaoUnitarias(gramatica)
	RemocaoInuteis(gramatica)

	fmt.Println("GRAMÁTICA SIMPLIFICADA!!! （っ＾▿＾）")
	fmt.Println()
}
