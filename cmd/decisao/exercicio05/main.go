package main

import "fmt"

func main() {
	var nome string
	var vendas, comissao float64
	fmt.Print("Nome do corretor: ")
	_, _ = fmt.Scan(&nome)
	fmt.Print("Valor das vendas: ")
	_, _ = fmt.Scan(&vendas)
	if vendas > 50_000 {
		comissao = vendas * 0.12
	} else if vendas >= 30_000 {
		comissao = vendas * 0.095
	} else {
		comissao = vendas * 0.07
	}
	fmt.Printf("%s receberá R$ %1.2f de comissão\n", nome, comissao)
}
