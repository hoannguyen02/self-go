package redis

import "github.com/go-redis/redis"

var Client *redis.Client

func CreateClient() {
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB: 0,
		Password: "",
	})
}