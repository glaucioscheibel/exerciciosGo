package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var leet string
	fmt.Print("Informe um frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()
	leet = strings.ToUpper(frase)
	leet = strings.ReplaceAll(leet, "D", "|)")
	leet = strings.ReplaceAll(leet, "I", "1")
	leet = strings.ReplaceAll(leet, "O", "0")
	leet = strings.ReplaceAll(leet, "A", "4")
	leet = strings.ReplaceAll(leet, "E", "3")
	leet = strings.ReplaceAll(leet, "B", "ß")
	fmt.Printf("L33t: %s\n", leet)
}
