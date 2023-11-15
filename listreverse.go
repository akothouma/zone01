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
	current := l.Head
	for current != nil {
		Next = current.Next
		current.Next = Prev
		Prev = current
		current = Next
	}
	l.Tail = Prev
	PrintValue(current)
}

func PrintValue(curr *NodeL) interface{} {
	return curr.Data
}
