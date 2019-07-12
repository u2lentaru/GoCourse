package main

import (
	"net/http"
	"time"
)

func main() {
	mirroredQuery()
}

func mirroredQuery() string {
	responses := make(chan string, 3)

	go func() {
		responses <- request("asia.site.io")
	}()

	go func() {
		responses <- request("europe.site.io")
	}()

	go func() {
		responses <- request("america.site.io")
	}()

	return <-responses
}

func request(hostname string) string {
	start := time.Now()
	http.Get(hostname)
	end := time.Now()
	return string(end.Sub(start))
}
