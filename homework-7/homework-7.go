package main

import (
	"fmt"
	"time"
)

func main() {
	//	a := 1
	//	if a == 0 {
	go spinner(50 * time.Millisecond)
	//	}

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
		for x := 0; x <= 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		y, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(y)
	}
}
