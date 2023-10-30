package test

import (
	"FormasNormais/helpers"
	"FormasNormais/helpers/gramatica"
	"fmt"
)

func ValidadeGreibachGramatica(gramatica *gramatica.Gramatica) {
	validacaoUnitario := false
	validacaoFormulaGeral := false

	for _, producoes := range gramatica.P {
		for _, regras := range producoes {
			if VerificarUnitario(regras, gramatica) {
				validacaoUnitario = true
			} else {
				if len(regras) == 1 {
					validacaoUnitario = false
					break
				}
			}

			if verificarPrimeiroSimboloUnitarioERestanteVariavel(regras, gramatica) {
				validacaoFormulaGeral = true
			} else {
				validacaoUnitario = false
				break
			}
		}
	}

	if validacaoUnitario && validacaoFormulaGeral {
		fmt.Print("\nTeste Gramática Greibach: Testes Passaram ✅✅ (っ＾▿＾)۶🍸🌟🍺٩(˘◡˘ )\n\n")
	} else {
		fmt.Print("\nTeste Gramática Greibach: Testes Reprovaram ❌❌ (╥﹏╥)\n\n")
	}
}

func verificarPrimeiroSimboloUnitarioERestanteVariavel(unitario []string, gramatica *gramatica.Gramatica) bool {
	if len(unitario) > 1 {
		if !helpers.VerificarSeEstarNasKeys(gramatica.Alf, unitario[0]) {
			return false
		}
	}

	for index, elm := range unitario {
		if !helpers.VerificarSeEstarNasKeys(gramatica.V, elm) && index > 0 {
			return false
		}
	}

	return true
}

func VerificarUnitario(unitario []string, gramatica *gramatica.Gramatica) bool {
	if len(unitario) == 1 {
		if helpers.VerificarSeEstarNasKeys(gramatica.Alf, unitario[0]) {
			return true
		}
	}

	return false
}


