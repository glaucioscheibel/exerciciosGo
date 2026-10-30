package main

import (
	"fmt"
	"strings"
)

func ehPrefixo(palavra, prefixo string) bool {
	return strings.HasPrefix(palavra, prefixo)
}

func main() {
	var palavra, prefixo string
	fmt.Print("Palavra: ")
	fmt.Scan(&palavra)
	fmt.Print("Prefixo: ")
	fmt.Scan(&prefixo)
	var msg string
	if !ehPrefixo(palavra, prefixo) {
		msg = " não"
	}
	fmt.Printf("%s%s é prefixo de %s\n", prefixo, msg, palavra)
}
