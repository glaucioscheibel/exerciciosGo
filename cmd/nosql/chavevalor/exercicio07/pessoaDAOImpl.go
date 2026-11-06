package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type PessoaDAOImpl struct {
	client *redis.Client
	ctx    context.Context
}

func NewPessoaDAOImpl() *PessoaDAOImpl {
	opt, _ := redis.ParseURL("redis://localhost:6379")
	return &PessoaDAOImpl{
		client: redis.NewClient(opt),
		ctx:    context.Background(),
	}
}

func (dao *PessoaDAOImpl) Create(p Pessoa) error {
	key := "contatos1:" + p.ID

	values := map[string]any{
		"id":       p.ID,
		"nome":     p.Nome,
		"telefone": p.Telefone,
		"email":    p.Email,
	}

	return dao.client.HSet(dao.ctx, key, values).Err()
}

func (dao *PessoaDAOImpl) Read(id string) (Pessoa, error) {
	key := "contatos1:" + id

	data, err := dao.client.HGetAll(dao.ctx, key).Result()
	if err != nil {
		return Pessoa{}, err
	}

	if len(data) == 0 {
		return Pessoa{}, fmt.Errorf("pessoa com id %s não encontrada", id)
	}

	return Pessoa{
		ID:       data["id"],
		Nome:     data["nome"],
		Telefone: data["telefone"],
		Email:    data["email"],
	}, nil
}

func (dao *PessoaDAOImpl) Update(p Pessoa) error {
	return dao.Create(p)
}

func (dao *PessoaDAOImpl) Delete(id string) error {
	key := "contatos:" + id
	return dao.client.Del(dao.ctx, key).Err()
}
