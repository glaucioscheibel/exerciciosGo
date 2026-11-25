package main

import "fmt"

type Pessoa struct {
	ID        int    `bson:"_id"`
	Nome      string `bson:"nome"`
	Sobrenome string `bson:"sobrenome"`
	Telefone  string `bson:"telefone"`
	Idade     int    `bson:"idade"`
}

func (p Pessoa) String() string {
	return fmt.Sprintf(
		"Pessoa [id=%d, nome=%s, sobrenome=%s, telefone=%s, idade=%d]",
		p.ID, p.Nome, p.Sobrenome, p.Telefone, p.Idade,
	)
}
