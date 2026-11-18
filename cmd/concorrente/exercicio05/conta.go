package main

import "sync"

type Conta struct {
	id    int
	saldo float64
	lock  sync.Mutex
}

func NewConta(id int, saldoInicial float64) *Conta {
	return &Conta{
		id:    id,
		saldo: saldoInicial,
	}
}

func (c *Conta) Deposito(valor float64) {
	c.saldo += valor
}

func (c *Conta) Saque(valor float64) {
	c.saldo -= valor
}

func (c *Conta) Saldo() float64 {
	return c.saldo
}

func (c *Conta) Lock() {
	c.lock.Lock()
}

func (c *Conta) Unlock() {
	c.lock.Unlock()
}
