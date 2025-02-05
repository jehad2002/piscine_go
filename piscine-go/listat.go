package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	f := l
	for i := 0; i < pos; i++ {
		if f == nil {
			return nil
		}
		f = f.Next
	}
	return f
}
