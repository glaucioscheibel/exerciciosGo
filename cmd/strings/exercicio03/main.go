package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Exercício 03 - Cálculo de dígito verificador")
	var codigo string
	scanner := bufio.NewScanner(os.Stdin)
	for ok := false; !ok; {
		fmt.Print("Informe um código de 5 algarismos: ")
		scanner.Scan()
		codigo = scanner.Text()
		ok = len(codigo) == 5
		if ok {
			for _, ch := range codigo {
				ok = unicode.IsDigit(ch)
				if !ok {
					fmt.Println("Código inválido. Digite apenas números.")
					break
				}
			}
		} else {
			fmt.Println("Código inválido. Digite exatamente 5 algarismos.")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao ler entrada:", err)
		return
	}
	a := int(codigo[0])
	b := int(codigo[1])
	c := int(codigo[2])
	d := int(codigo[3])
	e := int(codigo[4])
	digito := (6*a + 5*b + 4*c + 3*d + 2*e) % 7
	fmt.Printf("Código com dígito: %s-%d\n", codigo, digito)
}
