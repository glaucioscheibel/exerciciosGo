package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func soma(numeros []int16, ini int, qtde int, result chan int64) {
	var total int64
	fmt.Printf("somando de %d até %d\n", ini, ini+qtde)
	for i := ini; i < ini+qtde; i++ {
		total += int64(numeros[i])
	}
	result <- total
}

func executa(numeros []int16, qtdeThreads int, virtual bool, sequential int64) {
	hini := time.Now().UnixMilli()
	var wg sync.WaitGroup
	faixa := len(numeros) / qtdeThreads
	fmt.Printf("faixa %d qtdeThreads %d\n", faixa, qtdeThreads)
	inicio := make([]int, qtdeThreads)
	for i := range qtdeThreads {
		inicio[i] = i * faixa
	}
	channels := make([]chan int64, qtdeThreads)
	for i := range qtdeThreads {
		if virtual {
			wg.Go(func() {
				soma(numeros, inicio[i], faixa, channels[i])
			})
		} else {
			wg.Go(func() {
				runtime.LockOSThread()
				defer runtime.UnlockOSThread()
				soma(numeros, inicio[i], faixa, channels[i])
			})
		}
	}
	wg.Wait()
	var total int64
	for i := range qtdeThreads {
		total += <-channels[i]
	}
	millis := time.Now().UnixMilli() - hini
	speedup := float64(sequential) / float64(millis)
	eficiencia := speedup / float64(qtdeThreads*100)
	var tipo string
	if virtual {
		tipo = "Virtuais"
	} else {
		tipo = "Nativas"
	}
	fmt.Printf("Total %d Threads %s: %d\n", qtdeThreads, tipo, total)
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

	executa(numeros, 10, false, millis)
	executa(numeros, 100, false, millis)
	executa(numeros, cores, false, millis)
	executa(numeros, 10, true, millis)
	executa(numeros, 100, true, millis)
	executa(numeros, cores, true, millis)
}
