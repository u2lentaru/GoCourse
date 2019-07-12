package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mirroredQuery()
}

func mirroredQuery() string {
	responses := make(chan string, 3)

	go func() {
		responses <- request("golang.org")
	}()

	go func() {
		responses <- request("google.com")
	}()

	go func() {
		responses <- request("ya.ru")
	}()

	time.Sleep(10 * time.Second)

	//for tr := range responses {
	//	fmt.Println(tr)
	//}
	close(responses)

	return ""
}

func request(hostname string) string {
	start := time.Now()
	fmt.Println(hostname + " start " + start.String())
	http.Get(hostname)
	end := time.Now()
	fmt.Println(hostname + " end " + end.String())
	fmt.Println(hostname + " " + end.Sub(start).String())
	return hostname + end.Sub(start).String()
}
