package helpers

import "strings"

func LimparString(elemento string) string {
	return strings.TrimRight(elemento, "\x00")
}
