package main

import "fmt"

func main() {
	var digitos string
	fmt.Print("Informe dígitos: ")
	fmt.Scanln(&digitos)
	primeiro := true
	for _, c := range digitos {
		if primeiro {
			primeiro = false
		} else {
			fmt.Print(", ")
		}
		switch c {
		case '0':
			fmt.Print("zero")
		case '1':
			fmt.Print("um")
		case '2':
			fmt.Print("dois")
		case '3':
			fmt.Print("três")
		case '4':
			fmt.Print("quatro")
		case '5':
			fmt.Print("cinco")
		case '6':
			fmt.Print("seis")
		case '7':
			fmt.Print("sete")
		case '8':
			fmt.Print("oito")
		case '9':
			fmt.Print("nove")
		}
	}
	fmt.Println()
}
