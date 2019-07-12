package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	mirroredQuery()
}

func mirroredQuery() {
	var wg sync.WaitGroup
	responses := make(chan string, 3)

	go func() {
		wg.Add(1)
		defer wg.Done()
		responses <- request("golang.org")
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		responses <- request("google.com")
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		responses <- request("ya.ru")
	}()

	wg.Wait()
	//close(responses)
	time.Sleep(10 * time.Second)

	//for tr, ok := range responses {
	for {
		tr, ok := <-responses
		if !ok {
			break
		}
		fmt.Println(tr)
	}
	close(responses)

	return
}

func request(hostname string) string {
	start := time.Now()
	//fmt.Println(hostname + " start " + start.String())
	http.Get(hostname)
	end := time.Now()
	//fmt.Println(hostname + " end " + end.String())
	//fmt.Println(hostname + " " + end.Sub(start).String())
	return hostname + " response " + end.Sub(start).String()
}
