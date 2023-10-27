package helpers

import (
	"FormasNormais/helpers/gramatica"
	"fmt"
)

func PrintGramatica(gramatica *gramatica.Gramatica){
	PrintVariaveis(gramatica)
	PrintAlfabeto(gramatica)
	PrintProducoes(gramatica)
}

func PrintVariaveis(gramatica *gramatica.Gramatica){
	fmt.Print("Variaveis: ")
	for _,vars := range gramatica.V{
		fmt.Print(vars," ")
	}
	fmt.Println()
}

func PrintAlfabeto(gramatica *gramatica.Gramatica){
	fmt.Print("Alfabeto: ")
	for _,alf := range gramatica.Alf{
		fmt.Print(alf," ")
	}
	fmt.Println()
}

func PrintProducoes(gramatica *gramatica.Gramatica){
	fmt.Printf("Produções:\n")
	for chave,regras := range gramatica.P {
		fmt.Print(chave+": ")
		for _,regra := range regras{
			fmt.Print(regra)
		}
		fmt.Println()
	}
}