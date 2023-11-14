package piscine

type NodeL1 struct {
	Data interface{}
	Next *NodeL
}

type List2 struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List2) int {
	TempNode := l.Head
	count := 0
	for TempNode != nil {
		count++
		TempNode = TempNode.Next
	}

	return count
}
