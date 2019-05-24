package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)
	timers()
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

func timers() {
	tick := make(<-chan time.Time)    
	tick = time.Tick(1 * time.Second) 
	for countdown := 10; countdown > 0; countdown-- {
		moment := <-tick
		fmt.Println(moment.Format("15:04:05"), countdown)
	}
	fmt.Println("The rocket starts!")
}
