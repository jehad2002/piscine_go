package piscine

func Index(s string, toFind string) int {
	toFindLen := len(toFind)
	for i := 0; i <= len(s)-toFindLen; i++ {
		if s[i:i+toFindLen] == toFind {
			return i
		}
	}
	return -1
}
