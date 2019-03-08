package redis

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client
var once sync.Once

func init() {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
			PoolSize: 10,
		})

		pong, err := client.Ping().Result()
		fmt.Println(pong, err)
		// Output: PONG <nil>
	})
}

func NewRedisWithPool(poolSize int) *redis.Client {

	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
			PoolSize: poolSize,
		})

		pong, err := client.Ping().Result()
		fmt.Println(pong, err)
		// Output: PONG <nil>
	})

	return client
}

func NewRedis() *redis.Client {

	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
			PoolSize: 10,
		})

		pong, err := client.Ping().Result()
		fmt.Println(pong, err)
		// Output: PONG <nil>
	})

	return client
}

func redisOptions() *redis.Options {
	return &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB

	}
}
