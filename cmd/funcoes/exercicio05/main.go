package main

import "fmt"

func ehBissexto(ano int) bool {
	return (ano%4 == 0 && ano%100 != 0) || ano%400 == 0
}

func main() {
	var ano int
	fmt.Print("Ano: ")
	fmt.Scan(&ano)
	var msg string
	if ehBissexto(ano) {
		msg = " "
	} else {
		msg = " não "
	}
	fmt.Printf("O ano %d%sé bissexto\n", ano, msg)
}
