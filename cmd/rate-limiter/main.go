package main

import (
	"github.com/amskbs/rate-limiter.git/internal/calculator"
	"github.com/amskbs/rate-limiter.git/internal/generate"
	"github.com/amskbs/rate-limiter.git/internal/limiter"
)

func main() {
	calculator.Start()

	//lim := limiter.NewTokenBucketLimiter(10)
	lim := limiter.NewFixedWindowCounterLimiter(100)

	c := generate.New(0, 30)
	for {
		select {
		case <-c:
			allowed := lim.Allow(limiter.Task{})
			if allowed {
				calculator.Event()
			}
		}
	}
}
