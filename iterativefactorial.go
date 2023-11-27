package piscine

func IterativeFactorial(nb int) int {
	result := 1
	overflow := 1

	if nb < 0 {
		return 0
	}

	for i := 0; i < nb; i++ {
		r = r * overflow
		overflow++
		if i > 20 {
			return result
		}
	}

	return result
}
