package helpers

import (
	"FormasNormais/helpers/gramatica"
	"bufio"
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

	scanner := bufio.NewScanner(file)
	gramatica := &gramatica.Gramatica{
		P: make(map[string][][]string),
	}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "\t", "")

		if strings.Contains(line, "G:") || strings.Compare(line,"{") == 0 {
			continue
		}
		
		if strings.Contains(line, "V:") {
			line = strings.ReplaceAll(line, "V:[", "")
			line = strings.ReplaceAll(line, "],", "")
			gramatica.V = strings.Split(line, ",")
			continue
		}

		gramatica.V = limparStringAll(gramatica.V)

		if strings.Contains(line, "alf:") {
			line = strings.ReplaceAll(line, "alf:[", "")
			line = strings.ReplaceAll(line, "],", "")
			gramatica.Alf = strings.Split(line, ",")
			continue
		}

		gramatica.Alf = limparStringAll(gramatica.Alf)

		if strings.Contains(line, "P:{") {
			continue
		}
		char := rune(line[0])

		if char == 125 {
			continue
		}

		line = removeFirstNCharacters(line, 3)
		line = strings.ReplaceAll(line, "],", "")
		line = strings.ReplaceAll(line, "]", "")
		//talvez remover depois
		line = strings.ReplaceAll(line, "_", "")

		
		for _,element := range strings.Split(line, ",") {
			gramatica.P[string(char)] = append(gramatica.P[string(char)],strings.Split(element,";"))
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}

	return gramatica
}

func removeFirstNCharacters(s string, n int) string {
	if n >= len(s) {
		return ""
	}
	return s[n:]
}

func limparStringAll(elementos []string) []string {
	for _, elemento := range elementos {
		elemento = limparString(elemento)
	}

	return elementos
}

func limparString(elemento string) string {
	return strings.TrimRight(elemento, "\x00")
}
