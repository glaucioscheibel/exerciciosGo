package main

import (
	"fmt"
)

type Enxugador struct {
	limpos     *PilhaPratos
	escorredor *PilhaPratos
	nome       string
}

func NewEnxugador(nome string, limpos, escorredor *PilhaPratos) *Enxugador {
	return &Enxugador{
		nome:       nome,
		limpos:     limpos,
		escorredor: escorredor,
	}
}

func (e *Enxugador) Run() {
	defer fmt.Printf("%s terminou\n", e.nome)
	for {
		prato := e.escorredor.RetirarPrato()
		if prato == nil {
			break
		}
		e.enxugar(prato)
		e.limpos.ColocarPrato(prato)
	}
}

func (e *Enxugador) enxugar(prato *Prato) {
	prato.SetEstado(Limpo)
	fmt.Printf("%s sequei prato %d\n", e.nome, prato.id)
}
