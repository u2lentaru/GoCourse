package main

import "fmt"
import "math"

func main() {
	//    fmt.Println("Привет, мир!")
	Conv()
	triangleSquare()
}

// Выводит сумму в $
func Conv() {
	var rsum int
	const scrs int = 64
	fmt.Println("Сумма в рублях?")
	fmt.Scanln(&rsum)
	fmt.Println("Сумма $", rsum/scrs)
}

//Вычисляет площадь, периметр и гипотенузу прямоугольного треугольника.
func triangleSquare()
{
	var	leg1, leg2 float64 = 2,3
	var hyp, sq, perimeter float64
	hyp = math.Sqrt(math.Pow(leg1,2)+math.Pow(leg2,2))
	fmt.Println("Гипотенуза: ",hyp)
}