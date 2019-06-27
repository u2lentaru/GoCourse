package main

import (
	"calculator"
	"fmt"
	"math"
)

func main() {
	calcCall()
}

type perimeter interface {
	perimeter() float64
}

type square struct {
	edge float64
}

func (s square) perimeter() float64 {
	return 4 * s.edge
}

type circle struct {
	radius float64
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func calcCall() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "help" {
			fmt.Println("Calculate рассчитывает данные выражения, переданного в строке")
			continue
		}

		if input == "exit" {
			break
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
