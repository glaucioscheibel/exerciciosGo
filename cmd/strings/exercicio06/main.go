package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var iniciais strings.Builder
	fmt.Print("Informe o nome completo: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nome := scanner.Text()
	palavras := strings.SplitSeq(nome, " ")
	for palavra := range palavras {
		if palavra != "do" &&
			palavra != "dos" &&
			palavra != "da" &&
			palavra != "das" &&
			palavra != "de" &&
			palavra != "di" &&
			palavra != "du" &&
			palavra != "e" {
			iniciais.WriteString(string(palavra[0]))
		}
	}
	fmt.Printf("Iniciais: %s\n", iniciais.String())
}
