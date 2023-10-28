package formagreibach

import (
	"FormasNormais/helpers/gramatica"
	"FormasNormais/helpers/simplificacao"
	"fmt"
)

func Formaprechomsky(gramatica *gramatica.Gramatica){
	simplificacao.Simplificacao(gramatica)

	// RenomearVariaveis

	fmt.Printf("GRAMÁTICA NA FORMA PRÉ-CHOMSKY!!! ʕ•́ᴥ•̀ʔっ \n")
}

func renomearVariaveis(gramatica *gramatica.Gramatica){
	// renameMap := (make(map[string]string))

	for chave,valor := range gramatica.V {
		fmt.Println(chave,valor)
	}
}