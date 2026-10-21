package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var vogais int
	var espacos int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Informe uma frase: ")
	scanner.Scan()
	frase := scanner.Text()
	for i := 0; i < len(frase); i++ {
		c := string(frase[i])
		switch c {
		case " ":
			espacos++
		case "a", "e", "i", "o", "u":
			vogais++
		}
	}
	fmt.Printf("Na frase tem %d espaços e %d vogais.\n", espacos, vogais)
}
