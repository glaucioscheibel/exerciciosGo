package main

import (
	"fmt"
)

func main() {
	var estadoCivil string
	fmt.Print("Inicial do estado civil: ")
	fmt.Scan(&estadoCivil)
	estadoCivil = string(estadoCivil[0])
	switch estadoCivil {
	case "S":
		fmt.Println("Solteiro")
	case "C":
		fmt.Println("Casado")
	case "V":
		fmt.Println("Viúvo")
	case "D":
		fmt.Println("Divorciado")
	case "E":
		fmt.Println("dEsquitado")
	default:
		fmt.Println("Inicial inválida")
	}
}
