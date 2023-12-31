package simplificacao

import (
	"FormasNormais/helpers"
	"FormasNormais/helpers/gramatica"
	"fmt"
)

// precissa-se de refatoração
func RemocaoUnitarias(gramatica *gramatica.Gramatica) {
		// Achar Variaveis com unitarios
		comUnitarios := acharRegrasUnitarias(gramatica)
		
		//Corrigir  unitarios recursivos
		if(len(comUnitarios)>0){
			acharUnitariosRecursivos(comUnitarios,gramatica.P)
		}
		
		// remover unitarios
		if(len(comUnitarios)>0){
			removerUnitario(comUnitarios,gramatica)
		}

		if(len(comUnitarios)>0){
			RemocaoUnitarias(gramatica)
		}
}

func removerUnitario(variaveisComUnitarios []string,gramatica *gramatica.Gramatica){
	for _, variavesUnitarios := range variaveisComUnitarios{
		for index,regra := range gramatica.P[variavesUnitarios]{
			if(len(regra)==1 && helpers.IsVariavel(regra[0],gramatica.V)){
				gramatica.P[variavesUnitarios] = removerElementoArray(index,gramatica.P[variavesUnitarios])
				gramatica.P[variavesUnitarios] = adicionarRegas(gramatica.P[regra[0]],gramatica.P[variavesUnitarios])
			}
		}
	}
}

func adicionarRegas(adicionar [][] string, modificar [][] string) [][] string{
	for _,value := range adicionar {
		if(!verificarSeJaContem(modificar,value)){
			modificar = append(modificar, value)
		}
	}

	return modificar
}

func verificarSeJaContem(matriz [][]string, value []string) bool {
	
	encontrado := false

	for _, fatia := range matriz {
        if stringSlicesEqual(fatia, value) {
            encontrado = true
            break
        }
    }
	return encontrado
}

func stringSlicesEqual(slice1, slice2 []string) bool {
    if len(slice1) != len(slice2) {
        return false
    }

    for i := range slice1 {
        if slice1[i] != slice2[i] {
            return false
        }
    }

    return true
}

func acharUnitariosRecursivos(variaveisComUnitarios []string,regras map[string][][]string){
	for _, variavesUnitarios := range variaveisComUnitarios{
		el := regras[variavesUnitarios]
		regras[variavesUnitarios] = acharRecursao(variavesUnitarios,el)
	}
}

func acharRecursao(element string, regras [][]string)[][] string{
	for index,r := range regras{
		if (len(r) ==1) {
			if(r[0] == element){
				return removerElementoArray(index,regras)
			}
		}
	}

	return regras
}

func removerElementoArray(index int,array [][]string) [][]string {
	if index >= 0 && index < len(array) {
        // Use fatias para criar um novo array sem o elemento
        newArray := append(array[:index], array[index+1:]...)
		
		return newArray
    } else {
        fmt.Println("Índice fora do intervalo.")
    }
	return array
}

func acharRegrasUnitarias(gramatica *gramatica.Gramatica)[]string{
	var comUnitarios []string
	regras := gramatica.P
	for i,variaveis := range regras{
		for _,producoes := range variaveis{
			if len(producoes) == 1{
				if(helpers.InArray(producoes[0],gramatica.V) && !helpers.InArray(i,comUnitarios)){
					comUnitarios = append(comUnitarios, i)
				}
			}
		}
	}

	return comUnitarios
}

