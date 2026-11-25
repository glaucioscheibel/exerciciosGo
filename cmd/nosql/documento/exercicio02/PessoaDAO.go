package main

type PessoaDAO interface {
	Create(pessoa Pessoa) error
	Update(pessoa Pessoa) error
	Read(id int) (*Pessoa, error)
	Delete(id int) error
}
