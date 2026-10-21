package main

import "fmt"

func main() {
	var n int = -1
	var menor int
	var maior int
	var pares int
	var impares int
	var total int
	var qtde int
	for ok := true; ok; ok = (n != 0) {
		fmt.Print("Número: ")
		fmt.Scanf("%d", &n)
		if n != 0 {
			if qtde == 0 {
				menor = n
				maior = n
			} else {
				if n < menor {
					menor = n
				}
				if n > maior {
					maior = n
				}
			}
			if n%2 == 0 {
				pares++
			} else {
				impares++
			}
			qtde++
			total += n
		}
	}
	fmt.Printf("Maior: %d\n", maior)
	fmt.Printf("Menor: %d\n", menor)
	fmt.Printf("Soma: %d\n", total)
	fmt.Printf("Média: %1.2f\n", float64(total)/float64(qtde))
	fmt.Printf("Pares: %d\n", pares)
	fmt.Printf("Impares: %d\n", impares)
}
