package piscine

func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, i := range s {
		if !(i >= 'A' && i <= 'Z') {
			return false
		}
	}
	return true
}
