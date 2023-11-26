package main

import (
	"github.com/amskbs/rate-limiter.git/internal/limiter"
	"math/rand"
	"time"
)

func main() {
	//lim := limiter.NewTokenBucketLimiter(10)
	lim := limiter.NewFixedWindowCounterLimiter(1)
	for {
		time.Sleep(time.Duration(rand.Intn(20)+10) * time.Millisecond)
		_ = lim.TryEnqueue(limiter.Task{})
	}
}
