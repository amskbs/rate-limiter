package calculator

import (
	"fmt"
	"time"
)

var repo [5]int
var currentCell = 0

func Start() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("rate =  %d\n", repo[currentCell])
				currentCell++
				if currentCell == 5 {
					currentCell = 0
				}
				repo[currentCell] = 0
			}
		}
	}()
}

func Event() {
	repo[currentCell]++
}
