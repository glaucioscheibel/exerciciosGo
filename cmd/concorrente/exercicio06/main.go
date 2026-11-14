package main

import (
	"fmt"
	"sync"
)

const TotalPratos = 100

func main() {
	sujos := NewPilhaPratosSimples("Sujos", TotalPratos)
	limpos := NewPilhaPratosSimples("Limpos", TotalPratos)
	escorredor := NewPilhaPratos("Escorredor", 10, TotalPratos)

	for i := 1; i <= TotalPratos; i++ {
		sujos.ColocarPrato(NewPrato(i, Sujo))
	}

	fmt.Printf("Pilha com %d pratos sujos: %s\n", sujos.GetQuantidade(), sujos)
	fmt.Printf("Pilha com %d pratos molhados: %s\n", escorredor.GetQuantidade(), escorredor)
	fmt.Printf("Pilha com %d pratos limpos: %s\n", limpos.GetQuantidade(), limpos)

	var wg sync.WaitGroup

	lavador1 := NewLavador("Lavador-1", sujos, escorredor)
	enxugador1 := NewEnxugador("Enxugador-1", limpos, escorredor)
	enxugador2 := NewEnxugador("Enxugador-2", limpos, escorredor)

	wg.Go(func() {
		lavador1.Run()
	})
	wg.Go(func() {
		enxugador1.Run()
	})
	wg.Go(func() {
		enxugador2.Run()
	})

	wg.Wait()

	fmt.Printf("Pilha com %d pratos sujos: %s\n", sujos.GetQuantidade(), sujos)
	fmt.Printf("Pilha com %d pratos molhados: %s\n", escorredor.GetQuantidade(), escorredor)
	fmt.Printf("Pilha com %d pratos limpos: %s\n", limpos.GetQuantidade(), limpos)
}
