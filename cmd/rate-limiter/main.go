package main

import (
	"github.com/amskbs/rate-limiter.git/internal/calculator"
	"github.com/amskbs/rate-limiter.git/internal/generator"
	"github.com/amskbs/rate-limiter.git/internal/limiter"
)

func main() {
	calculator.Start()

	//lim := limiter.NewTokenBucketLimiter(10)
	lim := limiter.NewFixedWindowCounterLimiter(100)
	c := generator.New()
	for {
		select {
		case <-c:
			res := lim.TryEnqueue(limiter.Task{})
			if res.Enqueued {
				calculator.Event()
			}
		}
	}
}
