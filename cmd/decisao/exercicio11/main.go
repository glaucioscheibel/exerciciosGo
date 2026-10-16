package main

import "fmt"

func main() {
	var ano int
	var msg string
	fmt.Print("Ano: ")
	fmt.Scan(&ano)
	bissexto := (ano%4 == 0 && ano%100 != 0) || ano%400 == 0
	if bissexto {
		msg = " "
	} else {
		msg = " não "
	}
	fmt.Printf("O ano %d%sé bissexto\n", ano, msg)
}
