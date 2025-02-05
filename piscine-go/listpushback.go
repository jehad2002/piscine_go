package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNodel := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNodel
		l.Tail = newNodel
	} else {
		l.Tail.Next = newNodel
		l.Tail = newNodel
	}
}
