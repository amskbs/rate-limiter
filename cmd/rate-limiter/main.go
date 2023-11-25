package main

import (
	"github.com/amskbs/rate-limiter.git/internal/limiter"
)

func main() {
	lim := limiter.NewTokenBucketLimiter(1)
	for {
		_ = lim.TryEnqueue(limiter.Task{})
	}
}
