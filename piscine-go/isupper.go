package piscine

func IsUpper(s string) bool {
	for _, i := range s {
		if i < 65 || i > 90 {
			return false
		}
	}
	return true
}
