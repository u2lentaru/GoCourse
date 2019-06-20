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
	fmt.Println("Введите число для проверки чётности")
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
	fmt.Println("Введите число для проверки деления на 3")
	fmt.Scanln(&i)
	if i%3 == 0 {
		fmt.Println(i, "делится на 3 без остатка")
	} else {
		fmt.Println(i, "не делится на 3 без остатка")
	}
}

//Вывод первых 100 чисел последовательности Фибоначчи
func seqFibonacci100() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Всего:", r)
		}
	}()
	fmt.Println("1 число Фибоначчи 1")
	fnumb := 1
	fib(0, 1, &fnumb)
}

//рекурсия
func fib(f1, f2 float64, pfnumb *int) {
	f := f1
	f1 = f2
	f2 = f + f1
	*pfnumb++
	fmt.Println(*pfnumb, "число Фибоначчи", f2)
	if *pfnumb >= 100 {
		panic(*pfnumb)
	}
	fib(f1, f2, pfnumb)
}
