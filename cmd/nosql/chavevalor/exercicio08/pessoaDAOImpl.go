package main

import (
	"context"
	"encoding/json"

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
	key := "contatos2:" + p.ID
	jsonData, _ := json.Marshal(p)
	return dao.client.JSONSet(dao.ctx, key, "$", string(jsonData)).Err()
}

func (dao *PessoaDAOImpl) Read(id string) (Pessoa, error) {
	key := "contatos2:" + id

	jsonData, err := dao.client.JSONGet(dao.ctx, key).Result()
	if err != nil {
		return Pessoa{}, err
	}

	var p Pessoa
	err = json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		return Pessoa{}, err
	}
	return p, nil
}

func (dao *PessoaDAOImpl) Update(p Pessoa) error {
	return dao.Create(p)
}

func (dao *PessoaDAOImpl) Delete(id string) error {
	key := "contatos:" + id
	return dao.client.Del(dao.ctx, key).Err()
}
