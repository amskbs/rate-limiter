package limiter

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewTokenBucketRedisLimiter(rps int) Limiter {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	l := TokenBucketRedisLimiter{
		rdb:    rdb,
		rps:    rps,
		tokens: rps,
	}
	l.scheduleResetting()
	return &l
}

func (t *TokenBucketRedisLimiter) scheduleResetting() {
	result, err := t.rdb.Get(context.Background(), "tokens").Result()
	i, err := t.rdb.Exists(context.TODO(), "tokens").Result()

	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				t.resetTokens()
			}
		}
	}()
}

type TokenBucketRedisLimiter struct {
	rdb    *redis.Client
	rps    int
	tokens int
}

func (t *TokenBucketRedisLimiter) resetTokens() {
	t.tokens = t.rps
}

func (t *TokenBucketRedisLimiter) Allow() bool {
	if t.tokens == 0 {
		//fmt.Println("task has NOT been enqueued")
		return false
	}

	t.tokens--

	return true
}
