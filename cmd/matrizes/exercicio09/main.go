package main

import "fmt"

func main() {
	dist := [7][7]int{{0, 2, 11, 6, 15, 11, 1}, {2, 0, 7, 12, 4, 2, 15}, {11, 7, 0, 11, 8, 3, 13}, {6, 12, 11, 0, 10, 2, 1}, {15, 4, 8, 10, 0, 5, 13}, {11, 2, 3, 2, 5, 0, 14}, {1, 15, 13, 1, 13, 14, 0}}
	var origem, destino int
	fmt.Print("Cidade origem(1-7): ")
	fmt.Scan(&origem)
	fmt.Print("Cidade destino(1-7): ")
	fmt.Scan(&destino)
	fmt.Printf("Tempo da cidade de %d para %d é de %d horas\n\n", origem, destino, dist[origem-1][destino-1])
	fmt.Println("Informe as cidades da rota. 0 finaliza.")
	var cidades []int
	aux := -1
	for aux != 0 {
		fmt.Print("Cidade: ")
		fmt.Scan(&aux)
		if aux != 0 {
			cidades = append(cidades, aux)
		}
	}
	tempo := 0
	for i := 0; i < len(cidades)-1; i++ {
		tempo += dist[cidades[i]-1][cidades[i+1]-1]
	}
	fmt.Printf("Rota %v é de %d horas\n", cidades, tempo)
}
