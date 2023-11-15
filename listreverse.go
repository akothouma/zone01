package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	var Prev *NodeL
	var Next *NodeL
	curr := l.Head
	current := l.Head
	for current != nil {
		Next = current.Next
		current.Next = Prev
		Prev = current
		current = Next
	}
	l.Head = Prev
	l.Tail = curr
	PrintValue(current)
}

func PrintValue(curr *NodeL) interface{} {
	return curr.Data
}
