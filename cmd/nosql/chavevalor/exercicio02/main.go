package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func main() {
	opt, err := redis.ParseURL("redis://localhost:6379")
	ctx := context.Background()
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	defer client.Close()
	for i := range 10 {
		client.Set(ctx, "chave"+strconv.Itoa(i), i, 0)
	}
	chaves, _ := client.Keys(ctx, "chave*").Result()
	for _, chave := range chaves {
		valor, _ := client.Get(ctx, chave).Result()
		fmt.Printf("%s = %s\n", chave, valor)
	}
}
