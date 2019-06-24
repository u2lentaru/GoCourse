package main

import (
	"GoCourse/homework-3/queue"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	carsAndTrucksStructs()
	testingQueue()
	jsonAddressBook()
}

//Creating, modifying and comparing cars and trucks data structures
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

//map addressBook->json->file
func jsonAddressBook() {

	addressBook := make(map[string][]int)
	addressBook["Bob2"] = []int{89167243813}
	addressBook["Bob2"] = append(addressBook["Bob2"], 89155243628)

	b, err := json.Marshal(addressBook)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json for write in file", b)

	err = ioutil.WriteFile("AddressBook.json", b, 0644)
	if err != nil {
		fmt.Println(err)
	}

	rcontent, rerr := ioutil.ReadFile("AddressBook.json")
	if rerr != nil {
		fmt.Println(rerr)
	}

	fmt.Println("json read from file", rcontent)
	fmt.Println("string(json) read from file", string(rcontent))

}

/*
API server listening at: 127.0.0.1:40420
{Volvo 2014 16000} white
Checking a==b true
Checking a==c false
1
2
3

json for write in file [123 34 66 111 98 50 34 58 91 56 57 49 54 55 50 52 51 56 49 51 44 56 57 49 53 53 50 52 51 54 50 56 93 125]
json read from file [123 34 66 111 98 50 34 58 91 56 57 49 54 55 50 52 51 56 49 51 44 56 57 49 53 53 50 52 51 54 50 56 93 125]
string(json) read from file {"Bob2":[89167243813,89155243628]}
Process exiting with code: 0
*/
