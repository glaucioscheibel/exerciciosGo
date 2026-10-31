package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func soma(numeros []int16, ini int, qtde int, result chan int64) {
	var total int64
	for i := ini; i < ini+qtde; i++ {
		total += int64(numeros[i])
	}
	result <- total
}

func executa(numeros []int16, qtdeGoRoutines int, sequential int64) {
	hini := time.Now().UnixMilli()
	faixa := len(numeros) / qtdeGoRoutines
	inicio := make([]int, qtdeGoRoutines)
	for i := range qtdeGoRoutines {
		inicio[i] = i * faixa
	}
	channels := make([]chan int64, qtdeGoRoutines)
	for i := range channels {
		channels[i] = make(chan int64)
	}
	for i := range qtdeGoRoutines {
		go soma(numeros, inicio[i], faixa, channels[i])
	}
	var total int64
	for i := range qtdeGoRoutines {
		total += <-channels[i]
	}
	millis := time.Now().UnixMilli() - hini
	speedup := float64(sequential) / float64(millis)
	eficiencia := (speedup / float64(qtdeGoRoutines)) * 100
	fmt.Printf("Total %d GoRoutines: %d\n", qtdeGoRoutines, total)
	fmt.Printf("Tempo: %d milissegundos \n", millis)
	fmt.Printf("Speedup: %.2fx\n", speedup)
	fmt.Printf("Eficiência: %.2f%%\n\n", eficiencia)
}

func main() {
	cores := runtime.NumCPU()
	fmt.Printf("O seu ambiente de execução possui %d núcleos\n", cores)
	numeros := make([]int16, 1_000_000_000)
	for i := range numeros {
		numeros[i] = int16(rand.Intn(32_767))
	}

	hini := time.Now().UnixMilli()
	var soma int64
	for _, n := range numeros {
		soma += int64(n)
	}
	millis := time.Now().UnixMilli() - hini
	fmt.Printf("Soma sequencial: %d\n", soma)
	fmt.Printf("Tempo: %d milissegundos\n\n", millis)

	executa(numeros, 10, millis)
	executa(numeros, 100, millis)
	executa(numeros, cores, millis)
}
