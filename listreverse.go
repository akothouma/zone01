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
	current := l.Head
	for current != nil {
		l.Next = current.Next
		current.Next = l.Prev
		l.Prev = current
		current = l.Next
	}
	l.Head = l.Prev
}
