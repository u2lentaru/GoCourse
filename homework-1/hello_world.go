package main

import "fmt"

func main() {
	//    fmt.Println("Привет, мир!")
	Conv()
}

// Conv() выводит сумму в $
func Conv() {
	var rsum int
	const scrs int = 64
	fmt.Println("Сумма в рублях?")
	fmt.Scanln(&rsum)
	fmt.Println("rsum ", rsum, " scrs ", scrs)
	fmt.Println("Сумма $", rsum/scrs)
}
