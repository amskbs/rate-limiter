package calculator

import (
	"fmt"
	"time"
)

type Calculator struct {
	name        string
	repo        [5]int
	currentCell int
}

func New(name string) *Calculator {
	c := Calculator{
		name:        name,
		repo:        [5]int{},
		currentCell: 0,
	}

	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("[%s] rate =  %d\n", c.name, c.repo[c.currentCell])
				c.currentCell++
				if c.currentCell == 5 {
					c.currentCell = 0
				}
				c.repo[c.currentCell] = 0
			}
		}
	}()

	return &c
}

func (c *Calculator) Event() {
	c.repo[c.currentCell]++
}
