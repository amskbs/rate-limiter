package metrics

import (
	"fmt"
	"time"
)

type Reporter struct {
	name         string
	counters     [5]int
	currentIndex int
}

func NewReporter(name string) *Reporter {
	c := Reporter{
		name:         name,
		counters:     [5]int{},
		currentIndex: 0,
	}

	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("[%s] rate =  %d\n", c.name, c.counters[c.currentIndex])
				c.currentIndex++
				if c.currentIndex == 5 {
					c.currentIndex = 0
				}
				c.counters[c.currentIndex] = 0
			}
		}
	}()

	return &c
}

func (c *Reporter) Event() {
	c.counters[c.currentIndex]++
}
