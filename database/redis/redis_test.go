package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = NewRedisWithPool(20)
	client.FlushDB()
}

func Test_RedisSetnx(t *testing.T) {

}

func BenchmarkRedisSetNX(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		resp := client.SetNX("key"+string(i), 1, time.Second*50)
		lockSuccess, err := resp.Result()
		if err != nil {
			fmt.Println(err, "lock result: ", lockSuccess)
		}
	}
}

// func Benchamrk_RedisSetNX(b *testing.B) {
// 	b.ResetTimer()
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {

// 		}
// 	})
// }

func BenchmarkRedis(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			client.LPush("shzy", "value")
		}
	})
}

func BenchmarkSetExpire(b *testing.B) {

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := client.Set("key", "hello", 0).Err(); err != nil {
				b.Fatal(err)
			}
			if err := client.Expire("key", time.Second).Err(); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPipeline(b *testing.B) {
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.Pipelined(func(pipe redis.Pipeliner) error {
				pipe.Set("key", "hello", 0)
				pipe.Expire("key", time.Second)
				return nil
			})
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
