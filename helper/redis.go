package helper

import 	"github.com/redis/go-redis/v9"

var Client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

