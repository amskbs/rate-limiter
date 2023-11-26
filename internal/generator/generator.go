package generator

import (
	"github.com/amskbs/rate-limiter.git/internal/limiter"
	"math/rand"
	"time"
)

func New() chan interface{} {
	c := make(chan interface{})
	go func() {
		for {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			time.Sleep(time.Duration(r.Intn(20)) * time.Millisecond)
			c <- limiter.Task{}
		}
	}()
	return c
}
