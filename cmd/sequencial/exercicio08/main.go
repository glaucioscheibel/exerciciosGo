package main

import "fmt"

func main() {
	var tempoGasto float64
	var velocidadeMedia float64
	fmt.Print("Tempo gasto: ")
	fmt.Scan(&tempoGasto)
	fmt.Print("Velocidade média: ")
	fmt.Scan(&velocidadeMedia)
	distancia := tempoGasto * velocidadeMedia
	litrosUsados := distancia / 12
	fmt.Printf("Velocidade média: %.2f \n", velocidadeMedia)
	fmt.Printf("Tempo gasto: %.2f \n", tempoGasto)
	fmt.Printf("Distância: %.2f \n", distancia)
	fmt.Printf("Litros usados: %.2f \n", litrosUsados)
}
