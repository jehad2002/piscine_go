package piscine

type NodeL77 struct {
	Data interface{}
	Next *NodeL
}

type List77 struct {
	Head *NodeL
	Tail *NodeL
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	current := l.Head
	for current != nil {
		if cond(current) {
			f(current)
		}
		current = current.Next
	}
}

func StringToInt(node *NodeL) {
	if val, ok := node.Data.(int); ok {
		node.Data = val * 2
	}
}
