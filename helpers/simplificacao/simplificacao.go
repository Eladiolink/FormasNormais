package simplificacao

import (
	"FormasNormais/helpers/gramatica"
)

func Simplificacao(gramatica *gramatica.Gramatica) {
	// semLambda := RemocaoPalavraVazia(gramatica)
	// gramatica.Regras = semLambda

	RemocaoUnitarias(gramatica)
	// fmt.Println(semUnitaria)
}
