package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	mirroredQuery()
}

func mirroredQuery() {
	//var wg sync.WaitGroup
	responses := make(chan string, 3)

	go func() {
		responses <- request("http://geekbrains.ru")
	}()

	go func() {
		responses <- request("http://r0.ru")
	}()

	go func() {
		responses <- request("http://ya.ru")
	}()

	//time.Sleep(10 * time.Second)
	wg.Wait()

	for {
		tr, ok := <-responses
		if !ok {
			close(responses)
			break
		}

		fmt.Println(tr)
	}

	return
}

func request(hostname string) string {
	wg.Add(1)
	defer wg.Done()
	start := time.Now()
	_, err := http.Get(hostname)
	if err != nil {
		fmt.Println(err)
	}
	end := time.Now()
	return hostname + " response " + end.Sub(start).String()
}
