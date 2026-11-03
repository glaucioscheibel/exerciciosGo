package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

type safeList struct {
	lista []int
	mu    sync.Mutex
}

func insere(lista *safeList) {
	for true {
		lista.mu.Lock()
		lista.lista = append(lista.lista, rand.Intn(100))
		lista.mu.Unlock()
	}
}

func remove(lista *safeList) {
	for true {
		lista.mu.Lock()
		if len(lista.lista) > 0 {
			lista.lista = slices.Delete(lista.lista, 0, 1)
		}
		lista.mu.Unlock()
	}
}

func imprime(lista *safeList) {
	for true {
		lista.mu.Lock()
		fmt.Println("Lista atual:", lista.lista)
		lista.mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var lista safeList
	var wg sync.WaitGroup
	wg.Go(func() {
		insere(&lista)
	})
	wg.Go(func() {
		insere(&lista)
	})
	wg.Go(func() {
		remove(&lista)
	})
	wg.Go(func() {
		imprime(&lista)
	})
	wg.Wait()
}
