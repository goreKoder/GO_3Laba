// main.go
package main

import (
	"fmt"
	"mathutils" // Импортируйте ваш пакет
)

func main() {
	var n int
	fmt.Print("Введите неотрицательное целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	result, err := mathutils.Factorial(n)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Факториал %d равен %d\n", n, result)
}

//			go run zad_1/main.go mathutils.go
