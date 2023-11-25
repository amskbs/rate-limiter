package limiter

import (
	"fmt"
	"time"
)

func NewTokenBucketLimiter(rps int) Limiter {
	l := TokenBucketLimiter{
		rps:    rps,
		tokens: rps,
	}
	go refillEverySecond(&l)
	return &l
}

func refillEverySecond(l *TokenBucketLimiter) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			l.refillTokens()
		}
	}
}

type TokenBucketLimiter struct {
	rps    int
	tokens int
}

func (t *TokenBucketLimiter) refillTokens() {
	t.tokens = t.rps
}

func (t *TokenBucketLimiter) TryEnqueue(_ Task) EnqueueResult {
	if t.tokens == 0 {
		//fmt.Println("task has NOT been enqueued")
		return EnqueueResult{
			Enqueued: false,
		}
	}

	t.tokens--
	fmt.Printf("task enqueue time is %s\n", time.Now())

	return EnqueueResult{
		Enqueued: true,
	}
}
