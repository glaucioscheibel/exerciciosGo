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
	fmt.Println(client.Ping(ctx))
	fmt.Println(client.Echo(ctx, "Hello World!"))
}
