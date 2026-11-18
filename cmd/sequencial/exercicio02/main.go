package main

import "fmt"

func main() {
	var n1, n2, n3, n4 float32

	print("Digite a primeira nota: ")
	_, _ = fmt.Scanln(&n1)
	print("Digite a segunda nota: ")
	_, _ = fmt.Scanln(&n2)
	print("Digite a terceira nota: ")
	_, _ = fmt.Scanln(&n3)
	print("Digite a quarta nota: ")
	_, _ = fmt.Scanln(&n4)

	media := (n1 + n2 + n3 + n4) / 4

	fmt.Printf("A média das notas é: %.2f\n", media)
}
