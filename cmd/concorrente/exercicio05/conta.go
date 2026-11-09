package main

import "sync"

type conta struct {
	id    int
	saldo float64
	lock  sync.Mutex
}

func NewConta(id int, saldoInicial float64) *conta {
	return &conta{
		id:    id,
		saldo: saldoInicial,
	}
}

func (c *conta) Deposito(valor float64) {
	c.saldo += valor
}

func (c *conta) Saque(valor float64) {
	c.saldo -= valor
}

func (c *conta) Saldo() float64 {
	return c.saldo
}

func (c *conta) Lock() {
	c.lock.Lock()
}

func (c *conta) Unlock() {
	c.lock.Unlock()
}
