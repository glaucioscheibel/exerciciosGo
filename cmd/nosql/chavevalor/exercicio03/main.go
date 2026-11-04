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
	value, _ := client.Incr(ctx, "programa:execucoes").Result()
	fmt.Printf("execuções: %d\n", value)
}
