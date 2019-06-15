package main

import "fmt"

func main() {
	//    fmt.Println("Привет, мир!")
	Conv()
}

func Conv() {
	var rsum int8
	const scrs int8 = 64
	fmt.Println("Сумма в рублях?")
	fmt.Scanln(&rsum)
	fmt.Println("Сумма $", rsum/scrs)
}
