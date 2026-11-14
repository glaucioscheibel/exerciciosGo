package main

import (
	"fmt"
	"sync"
)

type PilhaPratos struct {
	mu              sync.Mutex
	cond            *sync.Cond
	nome            string
	pratos          []*Prato
	capacidade      int
	totalPratos     int
	pratosAcessados int
}

func NewPilhaPratos(nome string, capacidade, totalPratos int) *PilhaPratos {
	p := &PilhaPratos{
		nome:        nome,
		capacidade:  capacidade,
		totalPratos: totalPratos,
		pratos:      make([]*Prato, 0),
	}
	p.cond = sync.NewCond(&p.mu)
	return p
}

func NewPilhaPratosSimples(nome string, totalPratos int) *PilhaPratos {
	return NewPilhaPratos(nome, totalPratos, totalPratos)
}

func (p *PilhaPratos) ColocarPrato(prato *Prato) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.pratos) == p.capacidade {
		fmt.Printf("esperando para colocar prato %d no %s.\n", prato.id, p.nome)
		p.cond.Wait()
	}

	p.pratos = append(p.pratos, prato)
	p.pratosAcessados++
	fmt.Printf("Colocou o prato %d no %s. (%s: %d/%d)\n", prato.id, p.nome, p.nome, len(p.pratos), p.capacidade)

	p.cond.Broadcast()
}

func (p *PilhaPratos) RetirarPrato() *Prato {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.pratos) == 0 {
		if p.pratosAcessados >= p.totalPratos {
			return nil
		}
		fmt.Printf("esperando para retirar prato do %s.\n", p.nome)
		p.cond.Wait()
	}

	prato := p.pratos[0]
	p.pratos = p.pratos[1:] // Remove do início (comportamento de fila)

	fmt.Printf("Retirou o prato %d do %s. (%s: %d/%d)\n", prato.id, p.nome, p.nome, len(p.pratos), p.capacidade)

	p.cond.Broadcast()
	return prato
}

func (p *PilhaPratos) GetQuantidade() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.pratos)
}

func (p *PilhaPratos) String() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return fmt.Sprintf("%s:%v", p.nome, p.pratos)
}
