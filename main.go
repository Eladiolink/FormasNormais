package main

import (
	"FormasNormais/helpers"
	"fmt"
	"os"
)

func main() {
	// controller()
	d := helpers.File()
	for _, sub := range d.Regras["S"] {
		fmt.Println(sub)
	}
}

func controller() {
	switch os.Args[2] {
	case "C":
		fmt.Println("Forma Normal Chomsky")
	case "G":
		fmt.Println("Forma Normal Geibach")
	default:
		fmt.Println("Argunmento de entrada inválido!")
	}
}
