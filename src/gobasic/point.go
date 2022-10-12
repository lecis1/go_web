package gobasic

type functype func(int) bool

func Transferparameters(a int) int {
	a = a + 1
	return a
}

func Transferpoint(a *int) int {
	*a = *a + 1
	return *a
}

func IsOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func IsEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func Filter(slice []int, f functype) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
