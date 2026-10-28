package main

import "fmt"

func main() {
	notas := [5]float32{9.9, 9.7, 9.8, 10, 10}
	maior := notas[0]
	menor := notas[0]
	var soma float32
	for _, nota := range notas {
		if nota > maior {
			maior = nota
		}
		if nota < menor {
			menor = nota
		}
		soma += nota
	}
	media := (soma - maior - menor) / 3.0
	fmt.Printf("Média: %1.2f desconsiderando as notas %1.2f e %1.2f\n", media, menor, maior)
}
