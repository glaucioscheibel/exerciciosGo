package main

import "fmt"

func main() {
	fmt.Println("Pense em um número entre 1 e 1000, e tecle enter.")
	fmt.Scanln()
	inicio := 1
	fim := 1000
	acertei := false
	var resp string
	for !acertei {
		chute := (inicio + fim) / 2
		fmt.Printf("O número que você pensou é %d?\n", chute)
		fmt.Print("S-im, M-aior, N-menor: ")
		fmt.Scan(&resp)
		switch resp {
		case "S":
			acertei = true
		case "M":
			inicio = chute + 1
		case "N":
			fim = chute - 1
		default:
			fmt.Println("Resposta inválida.")
		}
	}
	fmt.Println("Acertei!!!")
}
