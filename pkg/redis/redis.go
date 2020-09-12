package redis

import "github.com/go-redis/redis"

func init()  {
	redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}