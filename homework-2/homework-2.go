package main

import "fmt"

func main() {
	parityChecking()
	divByTreeChecking()
	seqFibonacci100()
}

//Проверка чётности
func parityChecking() {
	var i int
	fmt.Println("Введите число для проверки")
	fmt.Scanln(&i)
	if i%2 == 0 {
		fmt.Println(i, "чётное")
	} else {
		fmt.Println(i, "нечётное")
	}
}

//Проверка деления на 3
func divByTreeChecking() {
	var i int
	fmt.Println("Введите число для проверки")
	fmt.Scanln(&i)
	if i%3 == 0 {
		fmt.Println("делится на 3 без остатка")
	} else {
		fmt.Println("не делится на 3 без остатка")
	}
}

//Вывод первых 100 чисел последовательности Фибоначчи
func seqFibonacci100() {
	//defer, panic, recover
}
