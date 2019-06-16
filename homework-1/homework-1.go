package main

import "fmt"
import "math"

func main() {
	currConv()
	triangleSquare()
	depositSum()
}

//Выводит сумму в $
func currConv() {
	var rsum int
	const scrs int = 64
	fmt.Println("Сумма в рублях?")
	fmt.Scanln(&rsum)
	fmt.Println("Сумма $", rsum/scrs)
}

//Вычисляет площадь, периметр и гипотенузу прямоугольного треугольника.
func triangleSquare() {
	var leg1, leg2 float64 = 3, 4
	var hyp, sq, perimeter float64
	hyp = math.Sqrt(math.Pow(leg1, 2) + math.Pow(leg2, 2))
	fmt.Println("Гипотенуза: ", hyp)
	sq = leg1 * leg2 / 2
	fmt.Println("Площадь: ", sq)
	perimeter = leg1 + leg2 + hyp
	fmt.Println("Периметр: ", perimeter)
}

//Вычисляет сумму вклада
func depositSum() {
	var startDep, endDep, percent float64
	fmt.Println("Начальная сумма вклада?")
	fmt.Scanln(&startDep)
	fmt.Println("Годовой процент?")
	fmt.Scanln(&percent)
	endDep = startDep * math.Pow(1+percent/100, 5)
	fmt.Println("Сумма вклада через 5 лет: ", endDep)
}
