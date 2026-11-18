package main

import "fmt"

func main() {
	var n1, n2, n3, n4, media, exame float32
	aprovado := false
	fmt.Print("Informe a nota1: ")
	_, _ = fmt.Scan(&n1)
	fmt.Print("Informe a nota2: ")
	_, _ = fmt.Scan(&n2)
	fmt.Print("Informe a nota3: ")
	_, _ = fmt.Scan(&n3)
	fmt.Print("Informe a nota4: ")
	_, _ = fmt.Scan(&n4)
	media = (n1 + n2 + n3 + n4) / 4
	if media >= 7 {
		aprovado = true
	} else {
		fmt.Printf("Media: %2.2f \n", media)
		fmt.Print("Nota do exame: ")
		_, _ = fmt.Scan(&exame)
		media = (media + exame) / 2
		if media >= 5 {
			aprovado = true
		}
	}
	fmt.Printf("Media final: %2.2f \n", media)
	if aprovado {
		fmt.Println("Aprovado!")
	} else {
		fmt.Println("Reprovado!")
	}
}
