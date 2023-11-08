package test

import (
	"FormasNormais/helpers"
	"FormasNormais/helpers/gramatica"
	"fmt"
)

func ValidadeChomskyGramatica(gramatica *gramatica.Gramatica){
	validacaoUnitario := false
	validacaoFormulaGeral := false
	condition := true
	
	for _, producoes := range gramatica.P {
		for _, regras := range producoes {
			if VerificarUnitario(regras, gramatica) {
				validacaoUnitario = true
			} else {
				if len(regras) == 1 {
					validacaoUnitario = false
					condition = false
					break
				}
			}

			if verificarRegras(regras, gramatica) {
				validacaoFormulaGeral = true
			} else {
				validacaoFormulaGeral = false
				condition = false
				break
			}
		}

		if !condition {
			break
		}
	}

	if validacaoUnitario && validacaoFormulaGeral {
		fmt.Print("\nTeste Gramática Chomsky: Testes Passaram ✅✅ (っ＾▿＾)۶🍸🌟🍺٩(˘◡˘ )\n\n")
	} else {
		fmt.Print("\nTeste Gramática Greibach: Testes Reprovaram ❌❌ (╥﹏╥)\n\n")
	}

}


func verificarRegras(regras []string ,gramatica * gramatica.Gramatica) bool {

			if len(regras)>2 {
				return false
			}

			if len(regras) == 2{
				if !helpers.InArray(regras[0],gramatica.V) && !helpers.InArray(regras[1],gramatica.V) {
					return false
				}
			}
	return true
}