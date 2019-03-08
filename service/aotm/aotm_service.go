package aotm

import (
	"sync"

	"github.com/shzy2012/testgin/database/redis"
)

type AotmService struct {
	LockKey string
}

func NewAotmService() *AotmService {
	return &AotmService{}
}

func (t *AotmService) AotmService() error {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			redis.Incr()
		}()
	}

	wg.Wait()
	return nil
}
