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
	fmt.Println("Possible chess knight moves from {3,3}", chessKnightMoves(3, 3))
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
		{"Mike", []int{89167245643}},
	}

	fmt.Println("Before sorting", u)
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

//structure for chessKnightMoves
type cbPoint struct {
	x int
	y int
}

func (cbp *cbPoint) setPoint(spx, spy int) {
	cbp.x = spx
	cbp.y = spy
}

func (cbp cbPoint) chessKnightPoints() []cbPoint {
	mycbp := make([]cbPoint, 0)
	var a cbPoint

	for i := -2; i <= 2; i++ {
		for j := -2; j <= 2; j++ {
			if i != j && i != -j && i != 0 && j != 0 {
				a.x = cbp.x + i
				a.y = cbp.y + j
				if a.x >= 1 && a.x <= 8 && a.y >= 1 && a.y <= 8 {
					mycbp = append(mycbp, a)
				}
			}
		}

	}

	return mycbp
}

//chessKnightMoves returns possible chess knight moves.
func chessKnightMoves(kx, ky int) []cbPoint {

	var ckp cbPoint
	ckp.setPoint(kx, ky)
	return ckp.chessKnightPoints()
}

//Hello from init
//> 5+2
//Result: 7
//> help
//Program calculates the data of the expression passed in the string. Type 'help' for help, 'exit' for exit.
//> exit
//Before sorting [{Bob [89167243814]} {Alex [89155243629 89155243630]} {Mike [89167245643]}]
//After sorting [{Alex [89155243629 89155243630]} {Bob [89167243814]} {Mike [89167245643]}]
//Possible chess knight moves from {3,3} [{1 2} {1 4} {2 1} {2 5} {4 1} {4 5} {5 2} {5 4}]
