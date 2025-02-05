package piscine

func Rot14(s string) string {
	h := []rune(s)
	for i, g := range h {
		if (g >= 'a' && g <= 'l') || (g >= 'A' && g <= 'L') {
			h[i] = g + 14
		}
		if (g >= 'm' && g <= 'z') || (g >= 'M' && g <= 'Z') {
			h[i] = g - 12
		}
	}
	return string(h)
}
