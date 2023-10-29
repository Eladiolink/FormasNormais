package export

import (
	"FormasNormais/helpers/gramatica"
	"fmt"
	"os"
)

func ExportToJson(gramatica *gramatica.Gramatica) {
	// Nome do arquivo a ser criado
	nomeArquivo := "out.json"

	// Cria o arquivo, se ele não existir, ou abre para escrita se existir
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer arquivo.Close()

	// Conteúdo a ser inserido no arquivo
	conteudo := "{\n"
	conteudo += "	G:{\n"
	conteudo = incluirVariaveis(conteudo, gramatica)
	conteudo = incluirAlfabeto(conteudo, gramatica)
	conteudo = incluirRegras(conteudo,gramatica)
	conteudo += "	}\n"
	conteudo += "}\n"

	// Escreve no arquivo
	_, err = arquivo.WriteString(conteudo)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Conteúdo inserido no arquivo com sucesso!")
}

func incluirVariaveis(conteudo string, gramatica *gramatica.Gramatica) string {
	conteudo += "		V: ["
	for i, variaveis := range gramatica.V {
		if i == 0 {
			conteudo += variaveis
			continue
		}
		conteudo += ", " + variaveis
	}
	conteudo += "],\n"
	return conteudo
}

func incluirAlfabeto(conteudo string, gramatica *gramatica.Gramatica) string {
	conteudo += "		alf: ["
	for i, variaveis := range gramatica.Alf {
		if i == 0 {
			conteudo += variaveis
			continue
		}

		conteudo += ", " + variaveis
	}
	conteudo += "],\n"
	return conteudo
}

func incluirRegras(conteudo string, gramatica *gramatica.Gramatica) string {
	conteudo += "		P: {\n"
	numeroTotal := len(gramatica.P)
	quantificador := 0
	for variavel,producao := range gramatica.P{
		quantificador+=1
		conteudo += "			"+variavel+": ["

		conteudo = incluirProducoes(conteudo,producao)

		if quantificador==numeroTotal{
			conteudo += "]\n"
		}else{
			conteudo += "],\n"
		}
		
	}

	conteudo += "		}\n"
	return conteudo
}

func incluirProducoes(conteudo string,producao [][]string) string {
	for index,regras := range producao{
		str := splitRegras(regras,";")

		if index == 0{
			conteudo+=str
			continue
		}

		conteudo +=", "+str
		
	}
	
	return conteudo
}

func splitRegras(regras []string,simbolo string) string {
	str := ""

	for index,elm := range regras{
		if index == 0{
			str += elm
			continue
		}
		
		str +=simbolo+elm
	}

	return str
}