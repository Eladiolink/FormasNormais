package simplificacao

import (
	"FormasNormais/helpers/gramatica"
)

func Simplificacao(gramatica *gramatica.Gramatica) {
	RemocaoPalavraVazia(gramatica)
	// gramatica.Regras = semLambda

	// fmt.Println(gramatica)
	// RemocaoUnitarias(gramatica)
	// RemocaoInuteis(gramatica)
	// fmt.Println(semUnitaria)
}
