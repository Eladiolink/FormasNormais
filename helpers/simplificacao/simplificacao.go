package simplificacao

import (
	"FormasNormais/helpers/gramatica"
	"fmt"
)

func Simplificacao(gramatica *gramatica.Gramatica) {
	// semLambda := RemocaoPalavraVazia(gramatica)
	// gramatica.Regras = semLambda

	semUnitaria := RemocaoUnitarias(gramatica)
	fmt.Println(semUnitaria)
}
