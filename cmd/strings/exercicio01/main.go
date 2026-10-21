package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var nomeInvertido strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Informe o seu nome:")
	scanner.Scan()
	nome := scanner.Text()
	for i := len(nome) - 1; i >= 0; i-- {
		nomeInvertido.WriteString(string(nome[i]))
	}
	fmt.Printf("Seu nome invertido é %s\n", nomeInvertido.String())
}
