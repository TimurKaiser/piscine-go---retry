package piscine

func IterativeFactorial(nb int) int {
	result := 1

	for i := 1; i <= nb; i++ {
		result = result * i
	}
	if i > 20 {
		return 0
	}
	return result
}
