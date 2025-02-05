package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)

	start := 0
	end := 0

	for end < len(str) {
		for end < len(str) && str[end] != ' ' {
			end++
		}

		word := str[start:end]

		if end < len(str) && str[end] == ' ' {
			end++
		}

		summary[word]++

		start = end
	}

	return summary
}
