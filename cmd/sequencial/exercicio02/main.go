package main

import "fmt"

func main() {
	var n1 float32
	var n2 float32
	var n3 float32
	var n4 float32

	var media float32

	print("Digite a primeira nota: ")
	fmt.Scanln(&n1)
	print("Digite a segunda nota: ")
	fmt.Scanln(&n2)
	print("Digite a terceira nota: ")
	fmt.Scanln(&n3)
	print("Digite a quarta nota: ")
	fmt.Scanln(&n4)

	media = (n1 + n2 + n3 + n4) / 4

	fmt.Printf("A média das notas é: %.2f\n", media)
}
