package piscine

func Fibonacci(index int) int {
	result := -1

	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	
	result = result + index
	return result
}
