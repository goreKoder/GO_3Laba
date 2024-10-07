package main

import (
	"fmt"
	// "GitHub/GO_3LABA/mathutils/factorial.go"
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

//			go run main.go zad_1/mathutils.go
//			C:\Users\gleby\Documents\GitHub\GO_3LABA\mathutils\factorial.go
//			go run main.go
