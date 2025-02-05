package piscine

type N struct {
	Data interface{}
	Next *NodeL
}

type Liste struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	length := 0
	n := l.Head
	for n != nil {
		length++
		n = n.Next
	}
	return length
}
