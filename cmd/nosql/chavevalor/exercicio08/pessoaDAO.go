package main

type PessoaDAO interface {
	Create(pessoa Pessoa) error
	Read(id string) (Pessoa, error)
	Update(pessoa Pessoa) error
	Delete(id string) error
}
