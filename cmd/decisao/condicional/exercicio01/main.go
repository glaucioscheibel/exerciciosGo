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
		break
	case "C":
		fmt.Println("Casado")
		break
	case "V":
		fmt.Println("Viúvo")
		break
	case "D":
		fmt.Println("Divorciado")
		break
	case "E":
		fmt.Println("dEsquitado")
		break
	default:
		fmt.Println("Inicial inválida")
		break
	}
}
