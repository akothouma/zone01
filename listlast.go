package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
	Name string
}

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	return l.Tail.Data
}
