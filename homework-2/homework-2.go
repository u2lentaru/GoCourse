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
		fmt.Println(i, "делится на 3 без остатка")
	} else {
		fmt.Println(i, "не делится на 3 без остатка")
	}
}

//Вывод первых 100 чисел последовательности Фибоначчи
func seqFibonacci100() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Println("1 число Фибоначчи 1")
	fmt.Println("2 число Фибоначчи 1")
	var fnumb uint64 = 2
	fib(1, 1, &fnumb)
}

//рекурсия
func fib(f1, f2 uint64, pfnumb *uint64) {
	f := f1
	f1 = f2
	f2 = f + f1
	fmt.Println(*pfnumb, "число Фибоначчи", f2)
	*pfnumb++
	if *pfnumb >= 100 {
		panic(*pfnumb)
	}
	fib(f1, f2, pfnumb)
}
