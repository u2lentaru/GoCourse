package main

import (
	"calculator"
	"fmt"
	"math"
	"sort"
)

func main() {
	calcCall()
	sortAddressBook()
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

type addressBook struct {
	abName string
	//	abPhones []int
	abPhones int
}

type utAddressBook []addressBook

func (u utAddressBook) Len() int {
	return len(u)
}

func (u utAddressBook) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u utAddressBook) Less(i, j int) bool {
	return u[i].abName < u[j].abName
}

//addressBook := make(map[string][]int)
//addressBook["Bob2"] = []int{89167243813}
//addressBook["Bob2"] = append(addressBook["Bob2"], 89155243628)

//func (a ByAge) Len() int           { return len(a) }
//func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func sortAddressBook() {
	u := utAddressBook{
		{"Bob", 89167243814},
		{"Alex", 89155243629},
	}

	fmt.Println(u)
	sort.Sort(utAddressBook(u))
	fmt.Println(u)
	//u[0].abName = "Alex"
	//u[0].abPhones = [89167243813]
	//u[1].abName = "Bob"
	//u[1].abPhones = [89155243628]

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
