package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}

	var result string

	if n%2 == 0 {
		result += "rock"
	}

	if n%3 == 0 {
		if result != "" {
			result += " and roll"
		} else {
			result += "roll"
		}
	}

	if result == "" {
		return "error: non divisible\n"
	}

	return result + "\n"
}
