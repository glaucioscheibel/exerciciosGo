package main

import "fmt"

func main() {
	var ladoA, ladoB, ladoC int
	var msg string
	fmt.Print("Lado A: ")
	_, _ = fmt.Scan(&ladoA)
	fmt.Print("Lado B: ")
	_, _ = fmt.Scan(&ladoB)
	fmt.Print("Lado C: ")
	_, _ = fmt.Scan(&ladoC)
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
	fmt.Println(msg)
}
