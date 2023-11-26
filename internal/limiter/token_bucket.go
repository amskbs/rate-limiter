package limiter

import (
	"time"
)

func NewTokenBucketLimiter(rps int) Limiter {
	l := TokenBucketLimiter{
		rps:    rps,
		tokens: rps,
	}
	l.scheduleResetting()
	return &l
}

func (t *TokenBucketLimiter) scheduleResetting() {
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

type TokenBucketLimiter struct {
	rps    int
	tokens int
}

func (t *TokenBucketLimiter) resetTokens() {
	t.tokens = t.rps
}

func (t *TokenBucketLimiter) Allow() bool {
	if t.tokens == 0 {
		//fmt.Println("task has NOT been enqueued")
		return false
	}

	t.tokens--

	return true
}
