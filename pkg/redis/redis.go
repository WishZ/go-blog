package redis

import (
	"go-blog/pkg/setting"
	"log"

	"github.com/go-redis/redis"
)

func init() {
	sec, err := setting.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	redis.NewClient(&redis.Options{
		Addr:     sec.Key("REDIS_HOST").String(),
		Password: sec.Key("REDIS_AUTH").String(), // no password set
		DB:       sec.Key("REDIS_DB").MustInt(1), // use default DB
	})
}
