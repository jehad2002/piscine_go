package piscine

func NRune(s string, n int) rune {
	x := '0'
	if n <= 0 || n > len(s) {
		return 0
	}
	for i, v := range s {
		if i == n-1 {
			x = v
		}
	}
	return x
}
