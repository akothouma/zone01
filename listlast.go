package piscine

type NodeL7 struct {
	Data interface{}
	Next *NodeL7
}

type List7 struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List7) interface{} {
	if l.Head == nil {
		return nil
	}
	return l.Tail.Data
}
