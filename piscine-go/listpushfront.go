package piscine

type n struct {
	Data interface{}
	Next *NodeL
}

type Lists struct {
	Head *n
	Tail *n
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
}
