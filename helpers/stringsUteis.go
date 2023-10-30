package helpers

import "strings"

func LimparString(elemento string) string {
	return strings.TrimRight(elemento, "\x00")
}

func InArray(element string,variaveis []string) bool {
	for _, elementos := range variaveis {
		if strings.Compare(LimparString(element),elementos) == 0 {
			return true
		}
	}

	return false
}

func VerificarSeEstarNasKeys(variaveis []string, elm string) bool {
	for _, key := range variaveis {
		if strings.Compare(LimparString(key), LimparString(elm)) == 0 {
			return true
		}
	}
	return false
}

func IsVariavel(value string, variaveis []string) bool {

	for _,key := range variaveis{
		if(key == value){
			return true
		}
	}
	return false
}