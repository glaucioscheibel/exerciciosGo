package main

import "fmt"

func tipoTriangulo(ladoA, ladoB, ladoC int) string {
	var msg string
	if ladoA < ladoB+ladoC && ladoB < ladoA+ladoC && ladoC < ladoA+ladoB {
		if ladoA == ladoB && ladoB == ladoC {
			msg = "Triângulo equilátero"
		} else if ladoA == ladoB || ladoB == ladoC || ladoC == ladoA {
			msg = "Triângulo isósceles"
		} else {
			msg = "Triângulo escaleno"
		}
	} else {
		msg = "Os lados não formam um triângulo"
	}
	return msg
}

func main() {
	var ladoA, ladoB, ladoC int
	fmt.Print("Lado A: ")
	fmt.Scan(&ladoA)
	fmt.Print("Lado B: ")
	fmt.Scan(&ladoB)
	fmt.Print("Lado C: ")
	fmt.Scan(&ladoC)
	fmt.Println(tipoTriangulo(ladoA, ladoB, ladoC))
}
