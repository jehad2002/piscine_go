package piscine

type NodeL1 struct {
	Data interface{}
	Next *NodeL
}

type Listis struct {
	Head *NodeL
	Tail *NodeL
}

func ListClear(l *List) {
	l.Head = nil
	l.Tail = nil
}
