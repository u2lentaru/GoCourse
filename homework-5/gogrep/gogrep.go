package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	query := os.Args[2]
	root := os.Args[1]
	found := 1

	fmt.Println("query", query)
	fmt.Println("root", root)
	//	fmt.Println("os.Args", os.Args)

	filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
		//		fmt.Println(path)
		if !file.IsDir() {
			file, err := os.Open(path)
			defer file.Close()
			if err != nil {
				return nil
			}
			scanner := bufio.NewScanner(file)
			for i := 1; scanner.Scan(); i++ {
				if strings.Contains(scanner.Text(), query) {
					found = 0
					fmt.Printf("%s/%s:%d: %s\n", root, path, i, scanner.Text())
				}
			}
		}
		return nil
	})

	/*	filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
			if !file.IsDir() {
				for i := 1; scanner.Scan(); i++ {
					if strings.Contains(scanner.Text(), query) {
						found = 0
						fmt.Printf("%s/%s:%d: %s\n", root, path, i, scanner.Text())
					}
			}
			return nil
		}
		}
	*/
}
