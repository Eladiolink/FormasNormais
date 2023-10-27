package simplificacao

import (
	"FormasNormais/helpers/gramatica"
	"fmt"
	"reflect"
)

//Precisa-se de refatoração

func RemocaoPalavraVazia(gramatica *gramatica.Gramatica) map[string][][]string {
	// verificar quais elementos tem palavra vazia
	var variaveisComPalavraVazia []string

	for key := range gramatica.P {
		if estaNoSlice(gramatica.P[key], "$") {
			variaveisComPalavraVazia = append(variaveisComPalavraVazia, key)
		}
	}

	// VERIFICAR NECESSIDADE
	variaveisComPalavraVazia = verificarRecusion(variaveisComPalavraVazia, gramatica)

	novaGramatica := make(map[string][][]string)

	for key := range gramatica.P {
		novaGramatica[key] = encontrarRemover(gramatica.P[key], "$")
	}

	for key := range gramatica.P {
		for _, variavel := range variaveisComPalavraVazia {
			for _, regra := range gramatica.P[key] {
				if regra[0] != "$" {
					novoElemento := removerElemento(regra, variavel)
					if !verificaOcorrencia(novaGramatica[key], novoElemento) {
						novaGramatica[key] = append(novaGramatica[key], novoElemento)
					}

				}
			}

			for _, regra := range novaGramatica[key] {
				if regra[0] != "$" {

					novoElemento := removerElemento(regra, variavel)

					if !verificaOcorrencia(novaGramatica[key], novoElemento) {
						novaGramatica[key] = append(novaGramatica[key], novoElemento)
					}

				}
			}
		}
	}
	
	return novaGramatica
}

func verificarRecusion(variaveisComPalavraVazia []string, gramatica *gramatica.Gramatica) []string {
	new := variaveisComPalavraVazia

	for key := range gramatica.P {
		if estaNoSlice2(gramatica.P[key], new) {
			if !inArray(key, new) {
				new = append(new, key)
			}
		}
	}

	index := len(new)

	if index != len(variaveisComPalavraVazia) {
		fmt.Println(new)
		return verificarRecusion(new, gramatica)
	}
	return new
}

func estaNoSlice(slice [][]string, value string) bool {

	for _, valor := range slice {
		for _, regras := range valor {
			if regras == value {
				return true
			}
		}
	}

	return false
}

func estaNoSlice2(slice [][]string, values []string) bool {

	for _, elements := range slice {
		if len(elements) == 1 && inArray(elements[0], values) {
			return true
		}
	}
	return false
}

func compareArrays(arr1, arr2 []string) bool {
	// Verificar se os tamanhos dos arrays são diferentes
	if len(arr1) != len(arr2) {
		return false
	}

	// Criar mapas para contar a frequência dos elementos em ambos os arrays
	countMap1 := make(map[string]int)
	countMap2 := make(map[string]int)

	for _, element := range arr1 {
		countMap1[element]++
	}

	for _, element := range arr2 {
		countMap2[element]++
	}

	// Comparar os mapas para verificar se os arrays têm os mesmos elementos
	return reflect.DeepEqual(countMap1, countMap2)
}

func encontrarRemover(slice [][]string, elemento string) [][]string {
	var new [][]string
	for _, value := range slice {
		if !valorInArray(value, elemento) {
			new = append(new, value)
		}
	}
	return new
}

func valorInArray(slice []string, elemento string) bool {
	for _, vallue := range slice {
		if vallue == elemento {
			return true
		}
	}

	return false
}

func verificaOcorrencia(slice [][]string, elemento []string) bool {
	for _, valor := range slice {
		if compareArrays(valor, elemento) {
			return true
		}
	}
	return false
}

func removerElemento(str []string, remover string) []string {
	var new []string

	for _, elm := range str {
		if elm != remover {
			new = append(new, elm)
		}
	}

	if len(new) == 0 {
		return str
	}

	return new
}
