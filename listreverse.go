package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
	Prev *NodeL
	Next *NodeL
}

func ListReverse(l *List) {
	l.Prev = nil
	l.Next = nil
	current := l.Head
	for current != nil {
		l.Next = current.Next
		current.Next = l.Prev
		l.Prev = current
		current = l.Next
	}
	current = l.Prev
	PrintValue(current)
}

func PrintValue(curr *NodeL) interface{} {
	return curr.Data
}
