package main

import (
	"fmt"
	"time"
)

func main()  {
	go sp(100 * time.Millisecond)
}

func sp(delay time.Duration)  {
	for {
		for _, r := range `-\|/`{
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}