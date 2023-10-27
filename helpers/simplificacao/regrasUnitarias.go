package simplificacao

import (
	"strings"
)


// precissa-se de refatoração
// func RemocaoUnitarias(gramatica *gramatica.Gramatica) map[string][]string {

// }

func copMap(original map[string][]string, copia map[string][]string) {
	for chave, valor := range original {
		copia[chave] = valor
	}
}

func add(regra []string, elements string, regras []string, unitario string, element string) []string {

	for _, index := range regras {
		regra = append(regra, index)
	}

	return regra
}

func remove(regra []string, element string, keys map[string]string) []string {
	var new []string

	for _, valor := range regra {
		if strings.Compare(limparString(element), limparString(element)) != 0 || !inArray(valor, keys) {
			new = append(new, valor)
		}
	}

	return new
}

func inArray(letra string, keys map[string]string) bool {

	for _, values := range keys {
		if strings.Compare(values, limparString(letra)) == 0 {
			return true
		}
	}

	return false
}

func getKeys(regras map[string][]string) map[string]string {
	mapa := make(map[string]string)
	for values := range regras {
		mapa[values] = values
	}

	return mapa
}

func acharUnitario(regras map[string][]string) []string {
	var unitarios []string
	var keys []string

	for key := range regras {
		keys = append(keys, key)
	}

	for key := range regras {
		for _, values := range regras[key] {
			values = limparString(values)
			if len(values) > 1 && !isKey(keys, values) {
				continue
			}

			if isArray(keys, values) {
				unitarios = append(unitarios, key)
			}
		}

	}

	return unitarios
}

func isKey(keys []string, element string) bool {
	for _, value := range keys {
		if strings.Compare(value, element) == 0 {
			return true
		}
	}

	return false
}

func limparString(elemento string) string {
	return strings.TrimRight(elemento, "\x00")
}

func isArray(keys []string, element string) bool {
	for _, values := range keys {
		if strings.Compare(values, element) == 0 {
			return true
		}
	}

	return false
}
