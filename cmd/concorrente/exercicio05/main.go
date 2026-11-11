package main

import (
	"math/rand"
	"sync"
	"time"
)

func movimento(b *banco, tempo time.Duration) {
	ini := time.Now()
	for time.Since(ini) < tempo {
		contaDebito := rand.Intn(100)
		contaCredito := rand.Intn(100)
		valor := rand.Float64() * 2500.0
		b.Transferencia(contaDebito, contaCredito, valor)
	}
}

func balanco(b *banco, tempo time.Duration) {
	ini := time.Now()
	for time.Since(ini) < tempo {
		b.ImprimeSaldos()
		time.Sleep(5 * time.Second)
	}
}

func main() {
	firmeza := NewBanco(100, 1000.0)
	var wg sync.WaitGroup
	for range 5 {
		wg.Go(func() {
			movimento(firmeza, 3*time.Minute)
		})
	}
	wg.Go(func() {
		balanco(firmeza, 3*time.Minute)
	})
	wg.Wait()
	firmeza.ImprimeSaldos()
}
