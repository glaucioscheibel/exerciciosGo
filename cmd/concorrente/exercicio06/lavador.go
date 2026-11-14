package main

import (
	"fmt"
)

type Lavador struct {
	sujos      *PilhaPratos
	escorredor *PilhaPratos
	nome       string
}

func NewLavador(nome string, sujos, escorredor *PilhaPratos) *Lavador {
	return &Lavador{
		nome:       nome,
		sujos:      sujos,
		escorredor: escorredor,
	}
}

func (l *Lavador) Run() {
	defer fmt.Printf("%s terminou\n", l.nome)
	for {
		prato := l.sujos.RetirarPrato()
		if prato == nil {
			break
		}
		l.lavar(prato)
		l.escorredor.ColocarPrato(prato)
	}
}

func (l *Lavador) lavar(prato *Prato) {
	prato.SetEstado(Molhado)
	fmt.Printf("%s lavou prato %d\n", l.nome, prato.id)
}
