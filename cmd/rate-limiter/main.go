package main

import (
	"github.com/amskbs/rate-limiter.git/internal/calculator"
	"github.com/amskbs/rate-limiter.git/internal/generate"
	"github.com/amskbs/rate-limiter.git/internal/limiter"
)

func main() {
	//lim := limiter.NewTokenBucketLimiter(10)
	//lim := limiter.NewFixedWindowCounterLimiter(100)
	lim := limiter.NewSlidingWindowLogLimiter(100)

	allowedRPSCalculator := calculator.New("ALLOWED")
	generatedRPSCalculator := calculator.New("GENERATED")
	c := generate.New(0, 20)
	for {
		select {
		case <-c:
			if lim.Allow() {
				allowedRPSCalculator.Event()
			}
			generatedRPSCalculator.Event()
		}
	}
}
