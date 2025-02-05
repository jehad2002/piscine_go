package piscine

func IsLower(s string) bool {
	for _, i := range s {
		if i < 97 || i > 122 {
			return false
		}
	}
	return true
}
