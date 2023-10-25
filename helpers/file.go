package helpers

import (
	"FormasNormais/helpers/gramatica"
	"fmt"
	"os"
	"strings"
)

func File() *gramatica.Gramatica {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return nil
	}
	defer file.Close()

	data := make([]byte, 100)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return nil
	}

	// Dividir a string em substrings usando quebras de linha como delimitador
	substrings := strings.FieldsFunc(string(data), func(r rune) bool {
		return r == '\n' || r == '\r'
	})

	//Remove Simbolos vazios
	for i, sub := range substrings {
		substrings[i] = strings.ReplaceAll(sub, " ", "")
	}

	myMap := make(map[string][]string)

	for _, sub := range substrings {
		variavel := strings.Split(sub, "->")
		regras := strings.Split(variavel[1], "|")
		myMap[variavel[0]] = regras
	}

	dados := &gramatica.Gramatica{}

	dados.Regras = myMap

	return dados
}
