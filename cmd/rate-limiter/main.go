package main

import (
	"github.com/amskbs/rate-limiter.git/internal/generate"
	"github.com/amskbs/rate-limiter.git/internal/limiter"
	"github.com/amskbs/rate-limiter.git/internal/metrics"
)

func main() {
	//lim := limiter.NewTokenBucketLimiter(10)
	//lim := limiter.NewFixedWindowCounterLimiter(100)
	lim := limiter.NewSlidingWindowLogLimiter(100)

	allowedReporter := metrics.NewReporter("ALLOWED")
	allReporter := metrics.NewReporter("ALL")
	c := generate.New(0, 20)
	for {
		select {
		case <-c:
			if lim.Allow() {
				allowedReporter.Event()
			}
			allReporter.Event()
		}
	}
}
