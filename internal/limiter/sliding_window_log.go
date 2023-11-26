package limiter

import (
	"sync/atomic"
	"time"
)

func NewSlidingWindowLogLimiter(rps int) Limiter {
	l := SlidingWindowLogLimiter{
		rps: int32(rps),
	}
	l.scheduleResetting()
	return &l
}

func (t *SlidingWindowLogLimiter) scheduleResetting() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				t.resetInterval()
			}
		}
	}()
}

type SlidingWindowLogLimiter struct {
	rps     int32
	counter atomic.Int32
}

func (t *SlidingWindowLogLimiter) resetInterval() {
	t.counter.Store(0)
}

func (t *SlidingWindowLogLimiter) Allow(_ Task) bool {
	if t.counter.Load() >= t.rps {
		//fmt.Println("task has NOT been enqueued")
		return false
	}

	t.counter.Add(1)

	return true
}
