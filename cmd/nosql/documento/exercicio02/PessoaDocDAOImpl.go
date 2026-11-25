package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type PessoaDocDAOImpl struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewPessoaDocDAOImpl(uri string) (*PessoaDocDAOImpl, error) {
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao MongoDB: %w", err)
	}

	collection := client.Database("contatos").Collection("pessoas")
	return &PessoaDocDAOImpl{
		client:     client,
		collection: collection,
	}, nil
}

func (dao *PessoaDocDAOImpl) Close(ctx context.Context) error {
	if dao.client == nil {
		return nil
	}
	return dao.client.Disconnect(ctx)
}

func (dao *PessoaDocDAOImpl) Create(pessoa Pessoa) error {
	ctx := context.Background()
	_, err := dao.collection.InsertOne(ctx, pessoa)
	if err != nil {
		return fmt.Errorf("erro ao inserir pessoa: %w", err)
	}
	return nil
}

func (dao *PessoaDocDAOImpl) Update(pessoa Pessoa) error {
	ctx := context.Background()
	filter := bson.M{"_id": pessoa.ID}
	update := bson.M{
		"$set": bson.M{
			"nome":      pessoa.Nome,
			"sobrenome": pessoa.Sobrenome,
			"telefone":  pessoa.Telefone,
			"idade":     pessoa.Idade,
		},
	}

	_, err := dao.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("erro ao atualizar pessoa: %w", err)
	}
	return nil
}

func (dao *PessoaDocDAOImpl) Read(id int) (*Pessoa, error) {
	ctx := context.Background()
	var p Pessoa

	err := dao.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&p)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("erro ao ler pessoa: %w", err)
	}
	return &p, nil
}

func (dao *PessoaDocDAOImpl) Delete(id int) error {
	ctx := context.Background()
	_, err := dao.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("erro ao remover pessoa: %w", err)
	}
	return nil
}
