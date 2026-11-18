package main

import "fmt"

func main() {
	var anos, meses, dias int
	fmt.Print("Anos: ")
	_, _ = fmt.Scan(&anos)
	fmt.Print("Meses: ")
	_, _ = fmt.Scan(&meses)
	fmt.Print("Dias: ")
	_, _ = fmt.Scan(&dias)
	dias += anos*365 + meses*30
	fmt.Printf("Você ja viveu %d dias\n", dias)
}
