package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {

	srcptr := flag.String("src", "", "source file")
	dstptr := flag.String("dst", "", "destination file")

	flag.Parse()

	fmt.Println("flags", *srcptr, *dstptr)

	content, err := ioutil.ReadFile(*srcptr)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(*dstptr, content, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
