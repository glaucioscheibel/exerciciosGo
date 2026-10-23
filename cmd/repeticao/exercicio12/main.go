package main

import "fmt"

func main() {
	paisA := 80_000
	paisB := 200_000
	anos := 0
	for paisA < paisB {
		anos++
		paisA += int(float64(paisA) * 0.03)
		paisB += int(float64(paisB) * 0.015)
	}
	fmt.Printf("O país A leva %d anos para ultrapassar o país B.\n", anos)
}
