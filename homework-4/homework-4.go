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

//New interface and several structures that satisfy it.

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

//Address book structure and implementation of interface Sort{}
type addressBook struct {
	abName   string
	abPhones []int
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

//Address book sorting
func sortAddressBook() {

	u := utAddressBook{
		{"Bob", []int{89167243814}},
		{"Alex", []int{89155243629, 89155243630}},
	}

	fmt.Println("Befor sorting", u)
	sort.Sort(u)
	fmt.Println("After sorting", u)
}

//Calculator calling. Calculator repository https://github.com/u2lentaru/Calculator
func calcCall() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "help" {
			fmt.Println("Program calculates the data of the expression passed in the string. Type 'help' for help, 'exit' for exit.")
			continue
		}

		if input == "exit" {
			break
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Result: %v\n", res)
		} else {
			fmt.Println("Could not calculate")
		}
	}
}

//Hello from init
//> 2+5
//Result: 7
//> help
//Program calculates the data of the expression passed in the string. Type 'help' for help, 'exit' for exit.
//> exit
//Befor sorting [{Bob [89167243814]} {Alex [89155243629 89155243630]}]
//After sorting [{Alex [89155243629 89155243630]} {Bob [89167243814]}]
