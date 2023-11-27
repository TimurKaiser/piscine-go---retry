package piscine

func Sqrt(nb int) int {
	result := 0
	if nb < 0 {
    	return 0
    }

	for result*result <= nb {
		if result*result == nb {
			return result
		}
		result++
	}

	return result
}
