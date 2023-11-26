package limiter

import (
	"container/list"
	"time"
)

func NewSlidingWindowLogLimiter(rps int) Limiter {
	return &SlidingWindowLogLimiter{
		rps: int32(rps),
		log: list.New(),
	}
}

type SlidingWindowLogLimiter struct {
	rps int32
	log *list.List
}

func (t *SlidingWindowLogLimiter) Allow() bool {
	for {
		e := t.log.Back()
		if e == nil {
			break
		}
		if e.Value.(int64) < time.Now().UnixNano()-int64(time.Second) {
			t.log.Remove(e)
		} else {
			break
		}
	}

	if t.log.Len() >= int(t.rps) {
		return false
	}

	t.log.PushFront(time.Now().UnixNano())

	return true
}
