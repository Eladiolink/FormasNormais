package simplificacao

import (
	"FormasNormais/helpers/gramatica"
	"fmt"
)

func Simplificacao(gramatica *gramatica.Gramatica) {
	semLambda := RemocaoPalavraVazia(gramatica)
	gramatica.Regras = semLambda

	fmt.Println(gramatica.Regras)
}
