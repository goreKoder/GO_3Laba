package mathutils

func Factorial(n int) int {

	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

//C:\Users\gleby\Documents\GitHub\GO_3LABA\mathutils\mathutils.go
