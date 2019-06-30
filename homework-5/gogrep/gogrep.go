package main

import (
	"fmt"
	"os"
)

func main() {
	query := os.Args[1]
	root := os.Args[2]

	//arg := os.Args[3]

	fmt.Println(query)
	fmt.Println(root)
	fmt.Println(os.Args)

	filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			for i := 1; scanner.Scan(); i++ {
				if strings.Contains(scanner.Text(), query) {
					found = 0
					fmt.Printf("%s/%s:%d: %s\n", root, path, i, scanner.Text())
				}
		}
		return nil
}
