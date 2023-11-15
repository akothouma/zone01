package piscine

type NodeL7 struct {
	Data interface{}
	Next *NodeL7
}

type List struct {
	Head *NodeL7
	Tail *NodeL7
}

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	return l.Tail.Data
}
