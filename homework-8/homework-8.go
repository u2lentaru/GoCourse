package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//run .\homework-8.go http://ya.ru http://r0.ru
func main() {
	serial()
}

func serial() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		serialFetch(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func serialFetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//bytes,err:=ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
