package formachomsky

import (
	formaprechomsky "FormasNormais/formasNormais/formaPreChomsky"
	"FormasNormais/helpers"
	"FormasNormais/helpers/gramatica"
	"FormasNormais/test"
	"fmt"
	"strconv"
)

type regraRemove struct {
	Key   string
	Regra []string
}

func FormaChomsky(gramatica *gramatica.Gramatica) *gramatica.Gramatica {
	formaprechomsky.Formaprechomsky(gramatica)
	helpers.PrintGramatica(gramatica)

	colocarNaquantidadeCorreta(gramatica, 0)

	test.ValidadeChomskyGramatica(gramatica)

	fmt.Printf("\nGRAMÁTICA NA FORMA CHOMSKY!!! (っ＾▿＾)۶🍸🌟🍺٩(˘◡˘ ) \n\n")

	helpers.PrintGramatica(gramatica)

	return gramatica
}

func colocarNaquantidadeCorreta(gramatica *gramatica.Gramatica, qt int) {
	var elementosRemover []regraRemove
	qtInicial := qt

	for keys, producoes := range gramatica.P {
		for indice, regras := range producoes {
			if len(regras) > 2 {
				qtInicial += 1
				elm := regraRemove{
					Key:   keys,
					Regra: regras,
				}

				elementosRemover = append(elementosRemover, elm)
				gramatica.P[keys] = removerElementoPorIndiceMatriz(gramatica.P[keys], indice)
			}
		}
	}

	index := 1
	for _, elm := range elementosRemover {
		parteUm, parteDois, erro := dividirArray(elm.Regra, 1)
		if erro == nil {
			novaRegra := []string{}
			novaRegra = append(novaRegra, parteUm[0])
			newVar := novaVar(gramatica)
			index += 1
			gramatica.P[newVar] = append(gramatica.P[newVar], parteDois)
			novaRegra = append(novaRegra, newVar)
			gramatica.P[elm.Key] = append(gramatica.P[elm.Key], novaRegra)

			gramatica.V = append(gramatica.V, newVar)
		}
	}

	if qtInicial != qt {
		colocarNaquantidadeCorreta(gramatica, qtInicial)
	}
}

func novaVar(gramatica *gramatica.Gramatica) string {
	var varZ []string
	for _, elm := range gramatica.V {
		firstElement := elm[0]
		if string(firstElement) == string("Z") {
			varZ = append(varZ, elm)
		}
	}

	return "Z" + strconv.Itoa(len(varZ)+1)
}

func dividirArray(arr []string, indice int) (primeiraParte []string, segundaParte []string, err error) {
	if indice < 0 || indice > len(arr) {
		err = fmt.Errorf("Índice fora dos limites do array")
		return
	}

	primeiraParte = arr[:indice]
	segundaParte = arr[indice:]

	return primeiraParte, segundaParte, nil
}

// Fatorar
func removerElementoPorIndiceMatriz(slice [][]string, indice int) [][]string {
	new := make([][]string, len(slice))
	if indice >= 0 && indice < len(slice) {
		// Remover o elemento pelo índice
		new = append(slice[:indice], slice[indice+1:]...)
	}

	return new
}
