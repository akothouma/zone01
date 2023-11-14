package piscine

type NodeL3 struct {
	Data interface{}
	Next *NodeL
}

type List3 struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List3) int {
	TempNode := l.Head
	count := 0
	for TempNode != nil {
		count++
		TempNode = TempNode.Next
	}

	return count
}
