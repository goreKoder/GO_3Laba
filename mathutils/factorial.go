package mathutils

import "fmt"

// Factorial вычисляет факториал заданного неотрицательного целого числа n.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("нельзя вычислить факториал отрицательного числа")
	}
	if n == 0 {
		return 1, nil
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result, nil
}
