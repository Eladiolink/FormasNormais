package test

import (
	"FormasNormais/helpers"
	"FormasNormais/helpers/gramatica"
	"fmt"
	"strings"
)

func ValidadeGreibachGramaticar(gramatica *gramatica.Gramatica) {
	validacaoUnitario := false
	validacaoFormulaGeral := false

	for _, producoes := range gramatica.P {
		for _, regras := range producoes {
			if verificarUnitario(regras, gramatica) {
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
		if !verificarSeEstarNasKeys(gramatica.Alf, unitario[0]) {
			return false
		}
	}

	for index, elm := range unitario {
		if !verificarSeEstarNasKeys(gramatica.V, elm) && index > 0 {
			return false
		}
	}

	return true
}

func verificarUnitario(unitario []string, gramatica *gramatica.Gramatica) bool {
	if len(unitario) == 1 {
		if verificarSeEstarNasKeys(gramatica.Alf, unitario[0]) {
			return true
		}
	}

	return false
}

func verificarSeEstarNasKeys(variaveis []string, elm string) bool {
	for _, key := range variaveis {
		if strings.Compare(helpers.LimparString(key), helpers.LimparString(elm)) == 0 {
			return true
		}
	}
	return false
}
