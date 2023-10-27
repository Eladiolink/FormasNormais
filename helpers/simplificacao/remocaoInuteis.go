package simplificacao

import (
	"FormasNormais/helpers/gramatica"
	"fmt"
	"strings"
)


func RemocaoInuteis(gramatica *gramatica.Gramatica){
	// fmt.Println(gramatica.P)
	// remover regras sem passo
	removerInalcancaveis(gramatica)
	removerRecursicoNotTerminal(gramatica)
}

func removerRecursicoNotTerminal(gramatica *gramatica.Gramatica){
	// encontar com produções finais
	comTerminais := encontarComProduçõesFinais(gramatica)
	// encontarComProduçõesFinais(gramatica)
	aRemover := slice1NotSlice2(gramatica.V,comTerminais)
	removerRegraEsimbolo(aRemover,gramatica)
}

func removerRegraEsimbolo(remover []string,gramatica *gramatica.Gramatica){
	for _,elm := range remover{
		delete(gramatica.P,elm)
		gramatica.V = removerElementoPorValor(gramatica.V,elm)
		for chave,regras := range gramatica.P {
			for i,regra := range regras {
				if inArray(elm,regra){
					gramatica.P[chave] = removerElementoPeloOIndice(regras,i)
				}
			}
		}
		
	}
}

func encontarComProduçõesFinais(gramatica *gramatica.Gramatica)[]string{
	var terminam []string
	for _, key := range gramatica.V {
		for _, producoes := range gramatica.P[key]{
			if(verificaTerminal(producoes,gramatica.V)){
				terminam = append(terminam, key)
			}
		}
	}

	terminam = encontarComProduçõesFinaisRecursion(gramatica,terminam,len(terminam))
	return terminam
}

func encontarComProduçõesFinaisRecursion(gramatica *gramatica.Gramatica,terminam []string, index int)[]string{
	newIndex := index

	for _, key := range gramatica.V {
		for _, producoes := range gramatica.P[key]{
			if(verificaTerminalRecusion(producoes,terminam) && !verificarSeJaEstar(terminam,key)){
				terminam = append(terminam, key)
			}
		}
	}

	newIndex = len(terminam)

	if(index != newIndex){
		terminam = encontarComProduçõesFinaisRecursion(gramatica,terminam,newIndex)
	}
	
	return terminam
}

func verificaTerminalRecusion(producoes []string, terminam []string) bool{
	for _,posicaoProducao := range producoes{
		if(verificarSeEstarNasKeys(terminam,posicaoProducao)){
			return true
		}
	}
	return false
}

func verificaTerminal(producoes []string,keys []string) bool{
	for _,posicaoProducao := range producoes{
		if(verificarSeEstarNasKeys(keys,posicaoProducao)){
			return false
		}
	}
	return true
}

func removerInalcancaveis(gramatica *gramatica.Gramatica){
	keys := gramatica.V
	valido := []string{"S"}
	for _,key := range keys{
		for _, regras := range gramatica.P[key]{
			for _, producoes := range regras {
				if(verificarSeEstarNasKeys(gramatica.V,producoes)){
					if(!verificarSeJaEstar(valido,producoes)){
						valido = append(valido, producoes)			
					}
				}
			}
		}
	}

	elm := slice1NotSlice2(gramatica.V,valido)

	for _,chaves := range elm{
		delete(gramatica.P,chaves)
		gramatica.V = removerElementoPorValor(gramatica.V,chaves)
	}

}

func slice1NotSlice2(slice1 []string, slice2 []string)[] string{
	elementosEmSlice2 := make(map[string]bool)

    for _, elemento := range slice2 {
        elementosEmSlice2[elemento] = true
    }

    elementosNaoEmSlice2 := []string{}

    for _, elemento := range slice1 {
        if _, ok := elementosEmSlice2[elemento]; !ok {
            elementosNaoEmSlice2 = append(elementosNaoEmSlice2, elemento)
        }
    }

	return elementosNaoEmSlice2
}

func verificarSeJaEstar(variaveis []string,elm string) bool {
	for _,key := range variaveis {
		if(strings.Compare(limparString(key), limparString(elm)) == 0){
			return true
		}
	}
	return false
}


func removerElementoPorValor(arr []string, value string) []string {
    var indexToRemove int
    found := false

    for i, v := range arr {
        if v == value {
            indexToRemove = i
            found = true
            break
        }
    }

    if found {
        arr = append(arr[:indexToRemove], arr[indexToRemove+1:]...)
    }

    return arr
}

func verificarSeEstarNasKeys(variaveis []string,elm string) bool {
	for _,key := range variaveis {
		if(strings.Compare(limparString(key), limparString(elm)) == 0){
			return true
		}
	}
	return false
}

func removerElementoPeloOIndice(arr [][]string, index int) [][]string {
    // Verifique se o índice está dentro dos limites do array
    if index < 0 || index >= len(arr) {
        fmt.Println("Índice fora dos limites")
        return arr
    }

    // Use a função append para criar um novo array de slices que exclui o slice no índice
    return append(arr[:index], arr[index+1:]...)
}