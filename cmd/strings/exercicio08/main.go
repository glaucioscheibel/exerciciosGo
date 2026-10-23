package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var palavra string
	var resultado strings.Builder
	fmt.Print("Informe uma palavra: ")
	fmt.Scanln(&palavra)
	for _, c := range palavra {
		if unicode.IsUpper(c) {
			resultado.WriteRune(unicode.ToLower(c))
		} else {
			resultado.WriteRune(unicode.ToUpper(c))
		}
	}
	fmt.Printf("Resultado: %s\n", resultado.String())
}
