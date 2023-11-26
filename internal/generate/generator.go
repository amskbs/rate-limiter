package generate

import (
	"github.com/amskbs/rate-limiter.git/internal/limiter"
	"math/rand"
	"time"
)

func New(minDelayMs, maxDelayMs int) chan interface{} {
	c := make(chan interface{})

	go func() {
		for {
			time.Sleep(time.Duration(random(minDelayMs, maxDelayMs)) * time.Millisecond)
			c <- limiter.Task{}
		}
	}()

	return c
}

func random(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}
