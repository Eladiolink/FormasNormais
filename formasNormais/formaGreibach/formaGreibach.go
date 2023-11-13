package formagreibach

import (
	formaprechomsky "FormasNormais/formasNormais/formaPreChomsky"
	"FormasNormais/helpers"
	"FormasNormais/helpers/gramatica"
	"FormasNormais/test"
	"fmt"

	// "os"
	"reflect"
	"strconv"
)

type elemento struct {
	Chave int
	Key   string
	Regra []string
}

type regraRemove struct {
	Key   string
	Regra []string
}

func FormaGreibach(gramatica *gramatica.Gramatica) *gramatica.Gramatica {
	formaprechomsky.Formaprechomsky(gramatica)

	// RenomearVariaveis
	renomearVariaveis(gramatica)

	// Verificar Ax -> Aj com x<j
	verificaVariaveisNumeros(gramatica, 0)

	// Remover Recursão a esquerda

	helpers.PrintProducoes(gramatica)
	gramatica = removerRecursaoEsquerda(gramatica, 0)
	helpers.PrintProducoes(gramatica)

	gramatica = relocRegras(gramatica, 0)

	fmt.Printf("\nGRAMÁTICA NA FORMA GREIBACH!!! (っ＾▿＾)۶🍸🌟🍺٩(˘◡˘ ) \n\n")

	test.ValidadeGreibachGramatica(gramatica)

	for key, elm := range gramatica.P {
		gramatica.P[key] = removeIdenticalSubSlices(elm)
	}

	helpers.PrintGramatica(gramatica)
	fmt.Println()

	return gramatica
}

// carece revisoes e melhorias :()
func relocRegras(gramatica *gramatica.Gramatica, qtInicial int) *gramatica.Gramatica {
	quantidade := qtInicial
	newGramatica := copiarGramatica(*gramatica)
	var elementosRemover []regraRemove

	for keys, producoes := range gramatica.P {
		for chave, regra := range producoes {
			if helpers.IsVariavel(regra[0], gramatica.V) {
				quantidade += 1

				res := adicionarRegasComSubstituicaoReturn(regra[0], newGramatica, regra, chave, keys)

				for _, newElement := range res {
					newGramatica.P[keys] = append(newGramatica.P[keys], newElement)
				}

				elm := regraRemove{
					Key:   keys,
					Regra: regra,
				}

				elementosRemover = append(elementosRemover, elm)

			}
		}
	}

	for _, elm := range elementosRemover {
		for index, regra := range newGramatica.P[elm.Key] {
			if compareString(regra, elm.Regra) {
				newGramatica.P[elm.Key] = removerElementoPorIndiceMatriz(newGramatica.P[elm.Key], index)
			}
		}
	}

	if quantidade != qtInicial {
		return relocRegras(newGramatica, 0)
	}

	return newGramatica
}

func verificarSeRegraEstaNaMatriz(matriz [][]string, regra []string) bool {

	for _, elm := range matriz {
		if compareString(elm, regra) {
			return true
		}
	}
	return false
}

func compareString(elm1 []string, elm2 []string) bool {
	if len(elm1) != len(elm2) {
		return false
	}

	for i, elm := range elm1 {
		if elm != elm2[i] {
			return false
		}
	}
	return true
}

func removerRecursaoEsquerda(gramatica *gramatica.Gramatica, qt int) *gramatica.Gramatica {
	// verificar recusão a esquerda
	qtInicial := qt
	var recursivo []elemento

	for variavel, producoes := range gramatica.P {
		for index, regras := range producoes {
			if variavel == regras[0] {
				qtInicial += 1
				elm := elemento{
					Chave: index,
					Key:   variavel,
					Regra: regras,
				}

				recursivo = append(recursivo, elm)
			}
		}
	}

	gramatica = tratarRecursaoEsquerda(gramatica, recursivo)

	if qtInicial != qt {
		gramatica = removerRecursaoEsquerda(gramatica, qtInicial)
	}

	return gramatica
}

func tratarRecursaoEsquerda(gramatica *gramatica.Gramatica, resursivos []elemento) *gramatica.Gramatica {
	newGramatica := copiarGramatica(*gramatica)

	for i, elm := range resursivos {
		fragmento := removerElementoPorIndice(newGramatica.P[elm.Key][elm.Chave], 0)
		newVar := "Z" + strconv.Itoa(i+1)
		newGramatica.P[newVar] = append(newGramatica.P[newVar], fragmento)
		newGramatica.V = append(newGramatica.V, newVar)
		fragmento = append(fragmento, newVar)
		newGramatica.P[newVar] = append(newGramatica.P[newVar], fragmento)

		for i, regra := range newGramatica.P[elm.Key] {
			if !verificarSeHaNoMap(resursivos, i) {
				regra = append(regra, newVar)
				newGramatica.P[elm.Key] = append(newGramatica.P[elm.Key], regra)
			}
		}
	}

	for _, elm := range resursivos {
		newGramatica.P[elm.Key] = removeElementoPeloElemento(newGramatica.P[elm.Key], elm.Regra)
	}

	return newGramatica
}

func removeElementoPeloElemento(matriz [][]string, elementoARemover []string) [][]string {
	var novaMatriz [][]string

	for _, linha := range matriz {
		encontrou := false

		// Verifica se a linha atual é igual ao elemento que queremos remover
		if len(linha) == len(elementoARemover) {
			encontrou = true
			for i := 0; i < len(linha); i++ {
				if linha[i] != elementoARemover[i] {
					encontrou = false
					break
				}
			}
		}

		// Se a linha atual não é o elemento a ser removido, adiciona à nova matriz
		if !encontrou {
			novaMatriz = append(novaMatriz, linha)
		}
	}

	return novaMatriz
}

func verificarSeHaNoMap(recursivos []elemento, i int) bool {
	for _, elm := range recursivos {
		if i == elm.Chave {
			return true
		}
	}

	return false
}

func renomearVariaveis(gramatica *gramatica.Gramatica) {
	renameMap := (make(map[string]string))

	for chave, valor := range gramatica.V {
		newVar := "A" + strconv.Itoa(chave+1)

		if !helpers.InArray(newVar, gramatica.V) {
			renameMap[valor] = newVar
		}

	}

	ranameCaracter := []string{}

	for i, _ := range renameMap {
		ranameCaracter = append(ranameCaracter, i)
	}

	for r, regras := range gramatica.P {
		for p, producoes := range regras {
			for c, caracter := range producoes {
				if helpers.InArray(caracter, ranameCaracter) {
					gramatica.P[r][p][c] = renameMap[caracter]
				}
			}
		}
	}

	for r, _ := range renameMap {
		gramatica.P[renameMap[r]] = gramatica.P[r]
		delete(gramatica.P, r)
	}

	var newVariaveis []string

	for _, elm := range gramatica.V {
		newVariaveis = append(newVariaveis, renameMap[elm])
	}

	for _, elm := range gramatica.V {
		if !helpers.InArray(elm, ranameCaracter) {
			newVariaveis = append(newVariaveis, elm)
		}
	}

	gramatica.V = newVariaveis

}

// Verificar  e Corrige Ax -> Aj com x<j
func verificaVariaveisNumeros(gramatica *gramatica.Gramatica, qtInicial int) {
	quantidade := qtInicial
	var elementosRemover []regraRemove
	newGramatica := copiarGramatica(*gramatica)

	for keys, regras := range gramatica.P {
		for chave, producoes := range regras {
			if len(producoes) > 1 {
				if keys > producoes[0] && helpers.IsVariavel(producoes[0], newGramatica.V) {
					quantidade += 1
					adicionarRegasComSubstituicao(producoes[0], newGramatica, producoes, chave, keys)

					elm := regraRemove{
						Key:   keys,
						Regra: producoes,
					}

					elementosRemover = append(elementosRemover, elm)
				}
			}
		}
	}

	for _, elm := range elementosRemover {
		for index, regra := range newGramatica.P[elm.Key] {
			if compareString(regra, elm.Regra) {
				newGramatica.P[elm.Key] = removerElementoPorIndiceMatriz(newGramatica.P[elm.Key], index)
			}
		}
	}

	if quantidade != qtInicial {
		verificaVariaveisNumeros(newGramatica, quantidade)
	}

	gramatica.P = newGramatica.P
}

func adicionarRegasComSubstituicaoReturn(adicionar string, gramatica *gramatica.Gramatica, dado []string, key int, chave string) [][]string {
	regraAdicionar := removerElementoPorIndice(gramatica.P[chave][key], 0)
	regra := [][]string{}

	for _, elm := range gramatica.P[adicionar] {
		joinRegra := make([]string, len(elm))
		copy(joinRegra, elm)

		for _, key := range regraAdicionar {
			joinRegra = append(joinRegra, key)
		}

		regra = append(regra, joinRegra)
	}

	return regra
}

func adicionarRegasComSubstituicao(adicionar string, gramatica *gramatica.Gramatica, dado []string, key int, chave string) {
	regraAdicionar := removerElementoPorIndice(gramatica.P[chave][key], 0)

	for _, elm := range gramatica.P[adicionar] {
		joinRegra := make([]string, len(elm))
		copy(joinRegra, elm)

		for _, key := range regraAdicionar {
			joinRegra = append(joinRegra, key)
		}
		gramatica.P[chave] = append(gramatica.P[chave], joinRegra)

	}
}

func removerElementoPorIndiceMatriz(slice [][]string, indice int) [][]string {
	new := make([][]string, len(slice))
	if indice >= 0 && indice < len(slice) {
		// Remover o elemento pelo índice
		new = append(slice[:indice], slice[indice+1:]...)
	}

	return new
}

func removerElementoPorIndice(slice []string, indice int) []string {
	novoSlice := make([]string, len(slice)-1)
	copy(novoSlice, slice[:indice])
	copy(novoSlice[indice:], slice[indice+1:])
	return novoSlice
}

func copiarMapa(original map[string][][]string) map[string][][]string {
	copia := make(map[string][][]string)

	for chave, valor := range original {
		copia[chave] = make([][]string, len(valor))
		for i, subSlice := range valor {
			copia[chave][i] = make([]string, len(subSlice))
			copy(copia[chave][i], subSlice)
		}
	}

	return copia
}

func copiarGramatica(original gramatica.Gramatica) *gramatica.Gramatica {
	// Criando uma cópia da estrutura original
	copia := gramatica.Gramatica{
		V:   make([]string, len(original.V)),
		Alf: make([]string, len(original.Alf)),
		P:   make(map[string][][]string),
	}

	// Copiando os slices de strings V e Alf
	copy(copia.V, original.V)
	copy(copia.Alf, original.Alf)

	// Copiando o mapa de slices de slices de strings P
	for chave, valor := range original.P {
		copia.P[chave] = make([][]string, len(valor))
		for i, subSlice := range valor {
			copia.P[chave][i] = make([]string, len(subSlice))
			copy(copia.P[chave][i], subSlice)
		}
	}

	return &copia
}

func removeIdenticalSubSlices(s [][]string) [][]string {
	var unique [][]string

	for _, subSlice := range s {
		found := false
		for _, existing := range unique {
			if reflect.DeepEqual(subSlice, existing) || equalSlices(subSlice, existing) {
				found = true
				break
			}
		}
		if !found {
			unique = append(unique, subSlice)
		}
	}
	return unique
}

func equalSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
