package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	opt, _ := redis.ParseURL("redis://localhost:6379")
	ctx := context.Background()
	client := redis.NewClient(opt)
	defer client.Close()
	nomes := []redis.Z{
		{Score: 1, Member: "Adalberto"},
		{Score: 2, Member: "Bianca"},
		{Score: 3, Member: "Celso"},
		{Score: 4, Member: "Diana"},
		{Score: 5, Member: "Edson"},
		{Score: 6, Member: "Fiona"},
		{Score: 7, Member: "Gilberto"},
		{Score: 8, Member: "Helena"},
		{Score: 9, Member: "Igor"},
		{Score: 10, Member: "Josefina"},
	}
	client.ZAdd(ctx, "nomes3", nomes...)
	tam, _ := client.ZCard(ctx, "nomes3").Result()
	values, _ := client.ZRange(ctx, "nomes3", 0, tam-1).Result()
	fmt.Printf("Tamanho da lista: %d\n", tam)
	fmt.Printf("Usuários na lista: %v\n", values)
}
