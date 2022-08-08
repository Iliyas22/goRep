package main

import (
	"fmt"
)

func main() {
	// (Задача1) Напишите программу для вычисления площади прямоугольника.
	// Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
	/* Task1
	var a, b int
	var s = 0
	fmt.Println("Input: a")
	fmt.Scanln(&a)
	fmt.Println("Input: b")
	fmt.Scanln(&b)
	s = a * b
	fmt.Println("Output: ", s)
	*/

	//task2
	//Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
	// Площадь круга должна вводиться пользователем с клавиатуры.
	/*
		var l, r, s, d float64
		fmt.Println("Введите площадь круга: ")
		fmt.Scanln(&s)
		r = math.Sqrt(math.Pi / s)
		d = 2 * r
		l = 2 * math.Pi * r
		fmt.Println("диаметр: ", d, "длинa: ", l)
	*/

	//task3
	//С клавиатуры вводится трехзначное число.
	// Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.

	var a int

	fmt.Println("Введите трехзначное число:")
	fmt.Scanln(&a)
	fmt.Println("Cотен: ", a/100)
	fmt.Println("Десятков: ", a/10%10)
	fmt.Println("Единиц: ", a%10)

}
