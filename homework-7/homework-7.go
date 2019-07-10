package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)

	//wait 10 seconds
	time.Sleep(10 * time.Second)

	//conveyor
	conveyor()
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

func conveyor() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
