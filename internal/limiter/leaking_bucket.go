package limiter

import (
	"time"
)

func NewLeakingBucketLimiter(rps int) Limiter {
	l := LeakingBucketLimiter{
		rps:    rps,
		tokens: rps,
	}
	go refillEverySecond(&l)
	return &l
}

func refillEverySecond(l *LeakingBucketLimiter) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			l.refillTokens()
		}
	}
}

type LeakingBucketLimiter struct {
	rps    int
	tokens int
}

func (t *LeakingBucketLimiter) refillTokens() {
	t.tokens = t.rps
}

func (t *LeakingBucketLimiter) Allow() bool {
	if t.tokens == 0 {
		//fmt.Println("task has NOT been enqueued")
		return false
	}

	t.tokens--

	return true
}
