package main

import "fmt"

func main() {
	var nome, sobrenome string
	fmt.Print("Informe nome: ")
	_, _ = fmt.Scan(&nome)
	fmt.Print("Informe sobrenome: ")
	_, _ = fmt.Scan(&sobrenome)
	nomeCompleto := nome + " " + sobrenome
	fmt.Printf("Seu nome completo é %s\n", nomeCompleto)
}
