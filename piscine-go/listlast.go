package piscine

type NodeL9 struct {
	Data interface{}
	Next *NodeL
}

type Listes struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	v := l.Tail

	if v == nil {
		return nil
	}

	return v.Data
}
