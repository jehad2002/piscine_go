package piscine

func Any(f func(string) bool, a []string) bool {
	IsNum := false
	for _, val := range a {
		if f(val) {
			IsNum = true
		}
	}
	return IsNum
}
