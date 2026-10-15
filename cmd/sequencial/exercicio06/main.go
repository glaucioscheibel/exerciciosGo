package main

import "fmt"

func main() {
	var nome string
	var sobrenome string
	var nomeCompleto string
	fmt.Print("Informe nome: ")
	fmt.Scan(&nome)
	fmt.Print("Informe sobrenome: ")
	fmt.Scan(&sobrenome)
	nomeCompleto = nome + " " + sobrenome
	fmt.Printf("Seu nome completo é %s\n", nomeCompleto)
}
