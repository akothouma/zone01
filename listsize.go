package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func Listsize(l *List) int {
	TempNode := l.Head
	count := 0
	for TempNode != nil {
		count++
		TempNode = TempNode.Next
	}

	return count
}
