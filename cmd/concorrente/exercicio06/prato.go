package main

import "fmt"

type Prato struct {
	id     int
	estado Estado
}

func NewPrato(id int, estado Estado) *Prato {
	return &Prato{
		id:     id,
		estado: estado,
	}
}

func (p *Prato) SetEstado(estado Estado) {
	p.estado = estado
}

func (p *Prato) GetId() int {
	return p.id
}

func (p *Prato) String() string {
	return fmt.Sprintf("%d:%s", p.id, p.estado)
}
