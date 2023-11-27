package piscine

func IterativeFactorial(nb int) int {
	result := 1
	overflow := 1

	if nb < 0 {
		return 0
	}

	for i := 0; i < nb; i++ {
		result = result * overflow
		overflow++
		if i > 20 {
			return 0
		}
	}

	return result
}
