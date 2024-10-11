package main //			go run main.go

import (
	"fmt"
	mathutils "test-vscode-module/zad_1"
	stringutils "test-vscode-module/zad_3"
)

func main() {
	var fac int
	fmt.Println("Задание 1-2:")
	fmt.Scan(&fac)
	fmt.Println(mathutils.Factorial(fac))

	var str string
	fmt.Println("Задание 3:")
	fmt.Scan(&str)
	fmt.Println(stringutils.Flipping(str))

	//задание 4
	fmt.Println("Задание 4:")
	var nam [5]int
	for i := 0; i < 5; i++ {
		nam[i] = i + 1
	}
	fmt.Println("Массив: ", nam)

	//задание 5
	var srez []int = nam[:]
	fmt.Println("Срез изначальный: ", srez)
	fmt.Println("Срез c доп. элементом: ", append(srez, 6))
	fmt.Println("Срез c доп. элементом: ", append(srez[2:], 6))

	//задание 6
	var string_srez []string = []string{"rwbveqb", "tebqfb", "favbfbvabb b", "ffbfb"}
	var nam_srez int
	var print_srez int
	for i := 0; i < len(string_srez); i++ {
		if nam_srez < len(string_srez[i]) {
			print_srez = i
		}
	}
	fmt.Println("Самая длинная строка: ", print_srez)
}

//			 go mod tidy
