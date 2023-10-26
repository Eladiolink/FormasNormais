package helpers

import (
	"FormasNormais/helpers/gramatica"
	"fmt"
	"strings"
	"reflect"
)

func Simplificacao(gramatica *gramatica.Gramatica) {
	semLambda := remocaoPalavraVazia(gramatica)
	gramatica.Regras = semLambda

	fmt.Println(gramatica.Regras)
}

func remocaoPalavraVazia(gramatica *gramatica.Gramatica) map[string][]string {
	// verificar quais elementos tem palavra vazia
	var variaveisComPalavraVazia []string

	for key := range gramatica.Regras {
		if estaNoSlice(gramatica.Regras[key], "Λ") {
			variaveisComPalavraVazia = append(variaveisComPalavraVazia, key)
		}
	}

	if(len(variaveisComPalavraVazia) == 0){
		return gramatica.Regras
	}

	for key := range gramatica.Regras {
		if estaNoSlice2(gramatica.Regras[key], variaveisComPalavraVazia) {
			variaveisComPalavraVazia = append(variaveisComPalavraVazia, key)
		}
	}

	fmt.Println(variaveisComPalavraVazia)
	novaGramatica := make(map[string][]string)

	for key := range gramatica.Regras {
		novaGramatica[key] = encontrarRemover(gramatica.Regras[key], "Λ")
	}

	// fmt.Println(novaGramatica)
	for key := range gramatica.Regras {
		for _, variavel := range variaveisComPalavraVazia {
			for _, regra := range gramatica.Regras[key] {
				if regra != "Λ" {
					novoElemento := strings.ReplaceAll(regra, variavel, "")
					if !verificaOcorrencia(novaGramatica[key], novoElemento) {
						novaGramatica[key] = append(novaGramatica[key], novoElemento)
					}

				}
			}

			for _, regra := range novaGramatica[key] {
				if regra != "Λ" {

					novoElemento := strings.ReplaceAll(regra, variavel, "")

					if !verificaOcorrencia(novaGramatica[key], novoElemento) {
						novaGramatica[key] = append(novaGramatica[key], novoElemento)
					}

				}
			}
		}
	}

	return novaGramatica
}

func estaNoSlice(slice []string, value string) bool {

	resultado := strings.Join(slice, " ")

	res := strings.Contains(resultado, value)

	strings.Split(resultado, " ")

	return res
}

func estaNoSlice2(slice []string, values []string) bool {

	for _, elements := range slice {
		splited := strings.Split(elements, "")
		return compareArrays(splited, values)
	}
	return true
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

func encontrarRemover(slice []string, elemento string) []string {
	for i, value := range slice {
		if strings.Compare(value, elemento) >= 0 {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice // Element not found
}

func verificaOcorrencia(slice []string, elemento string) bool {
	for _, valor := range slice {
		if valor == elemento {
			return true
		}
	}
	return false
}
