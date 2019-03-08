package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func Incr() {

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	fmt.Println(err, "lock result: ", lockSuccess)

	cntString, err := client.Get(counterKey).Result()
	cntValue, _ := strconv.Atoi(cntString)

	if err != nil && err != redis.Nil {
		panic(err)

	} else {

		cntValue++
		println("current counter is ", cntValue)
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()

		if err != nil {
			// log err
			println("set value error!")
		}
	}

	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}
