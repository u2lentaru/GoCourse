package main

import "fmt"
import "GoCourse/homework-3/queue"

func main() {
	carsAndTrucksStructs()
	testingQueue()
}

//Creating, modification and comparing cars and trucks data structures
func carsAndTrucksStructs() {

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

//Testing queue (FIFO)
func testingQueue() {

	queue.Push("1")
	queue.Push("2")

	fmt.Println(queue.Pop())

	queue.Push("3")

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}

/*
{Volvo 2014 16000} white
Checking a==b true
Checking a==c false
1
2
3

Process exiting with code: 0
*/
