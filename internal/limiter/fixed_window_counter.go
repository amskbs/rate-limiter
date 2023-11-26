package limiter

import (
	"sync/atomic"
	"time"
)

func NewFixedWindowCounterLimiter(rps int) Limiter {
	l := FixedWindowCounterLimiter{
		rps: int32(rps),
	}
	l.scheduleResetting()
	return &l
}

func (t *FixedWindowCounterLimiter) scheduleResetting() {
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

type FixedWindowCounterLimiter struct {
	rps     int32
	counter atomic.Int32
}

func (t *FixedWindowCounterLimiter) resetInterval() {
	t.counter.Store(0)
}

func (t *FixedWindowCounterLimiter) Allow() bool {
	if t.counter.Load() >= t.rps {
		//fmt.Println("task has NOT been enqueued")
		return false
	}

	t.counter.Add(1)

	return true
}
