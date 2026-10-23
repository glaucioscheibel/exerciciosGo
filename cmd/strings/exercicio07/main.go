package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra1, palavra2, ehAnagrama string
	fmt.Print("Palavra 1: ")
	fmt.Scanln(&palavra1)
	fmt.Print("Palavra 2: ")
	fmt.Scanln(&palavra2)
	anagrama := len(palavra1) == len(palavra2)
	if anagrama {
		aux := palavra2
		for _, c := range palavra1 {
			aux = strings.Replace(aux, string(c), "", 1)
		}
		anagrama = len(aux) == 0
	}
	if anagrama {
		ehAnagrama = ""
	} else {
		ehAnagrama = " não"
	}
	fmt.Printf("%s e %s%s são anagramas\n", palavra1, palavra2, ehAnagrama)
}
