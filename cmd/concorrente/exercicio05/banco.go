package main

import (
	"sync/atomic"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type banco struct {
	contas            []*conta
	transacoes        atomic.Int64
	transacoesDoSaldo int64
}

func NewBanco(qtdeContas int, saldoInicial float64) *banco {
	b := &banco{}
	b.transacoes.Store(0)
	b.contas = make([]*conta, qtdeContas)
	for i := range qtdeContas {
		b.contas[i] = NewConta(i, saldoInicial)
	}
	return b
}

func (b *banco) Transferencia(contaDebito, contaCredito int, valor float64) {
	if contaDebito == contaCredito {
		return
	}
	menor := min(contaDebito, contaCredito)
	maior := max(contaDebito, contaCredito)
	b.contas[menor].Lock()
	b.contas[maior].Lock()
	b.contas[contaDebito].Saque(valor)
	b.contas[contaCredito].Deposito(valor)
	b.contas[menor].Unlock()
	b.contas[maior].Unlock()
	b.transacoes.Add(1)
}

func (b *banco) getSaldoTotal() float64 {
	var total float64
	for _, conta := range b.contas {
		conta.Lock()
	}
	b.transacoesDoSaldo = b.transacoes.Load()
	for _, conta := range b.contas {
		total += conta.Saldo()
		conta.Unlock()
	}
	return total
}

func (b *banco) ImprimeSaldos() {
	p := message.NewPrinter(language.BrazilianPortuguese)
	p.Printf("Saldo do banco: %.2f Transações: %.d\n", b.getSaldoTotal(), b.transacoesDoSaldo)
}
