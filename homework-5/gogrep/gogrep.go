package main

import (
	"fmt"
	"os"
)

func main() {
	query := os.Args[0]
	root := os.Args[1]

	//arg := os.Args[3]

	fmt.Println(query)
	fmt.Println(root)
	fmt.Println(os.Args)
}
