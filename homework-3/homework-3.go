package main

import "fmt"

func main() {
	carsAndTrucksStracts()
}

//Creating, modification and comparing cars and trucks data structures
func carsAndTrucksStracts() {

	type car struct {
		model string
		year  int
		color string
	}

	type truck struct {
		model        string
		year         int
		loadCapacity float32
	}

	a := car{"Ford Focus", 2014, "white"}
	b := car{year: 2014, model: "Ford Focus", color: "white"}
	c := car{"Toyota Camry", 2012, "black"}
	t := truck{"Volvo", 2014, 12000}

	t.loadCapacity = 16000
	fmt.Println(t, a.color)
	fmt.Println("Checking a==b", a == b)
	fmt.Println("Checking a==c", a == c)
}
