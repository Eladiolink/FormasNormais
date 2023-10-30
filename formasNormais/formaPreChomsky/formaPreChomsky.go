package formaprechomsky

import (
	"FormasNormais/helpers"
	"FormasNormais/helpers/gramatica"
	"FormasNormais/helpers/simplificacao"
	"fmt"
)


func Formaprechomsky(gramatica *gramatica.Gramatica){
	simplificacao.Simplificacao(gramatica)

	// Criar Novas Váriaveis
	novasVariaveisComTerminais(gramatica)

	fmt.Printf("GRAMÁTICA NA FORMA PRÉ-CHOMSKY!!! ʕ•́ᴥ•̀ʔっ \n")
}

func novasVariaveisComTerminais(gramatica *gramatica.Gramatica){
	var novasVariaveis []string
	mapa := (make(map[string]string))

	for chave,elementos := range gramatica.P {
		for i,producoes := range elementos{
			if len(producoes)>1 {
				for j,caracter := range producoes{
					if (!helpers.IsVariavel(caracter,gramatica.V)){
						novaVariavel := "V"+caracter
						if(!helpers.IsVariavel(novaVariavel,novasVariaveis)){
							mapa[novaVariavel] = caracter
							el := []string{caracter}
							novasVariaveis = append(novasVariaveis, novaVariavel)
							gramatica.P[novaVariavel] = append(gramatica.P[novaVariavel],el)
							gramatica.V = append(gramatica.V, novaVariavel)
						}

						gramatica.P[chave][i][j] = novaVariavel
					}
				}
			}
		}
	}
}

