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
	nomes := []string{"Adalberto", "Bianca", "Celso", "Diana", "Edson", "Fiona", "Gilberto", "Helena", "Igor", "Josefina"}
	client.RPush(ctx, "nomes1", nomes)
	tam, _ := client.LLen(ctx, "nomes1").Result()
	values, _ := client.LRange(ctx, "nomes1", 0, -1).Result()
	fmt.Printf("Tamanho da lista: %d\n", tam)
	fmt.Printf("Usuários na lista: %v\n", values)
}
