package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head     *NodeL
	Tail     *NodeL
	PrevNode *NodeL
	NextNode *NodeL
}

func ListReverse(l *List) {
	l.PrevNode = nil
	l.NextNode = nil
	for l.Head != nil {
		l.NextNode = l.Head.Next
		l.Head.Next = l.PrevNode
		l.PrevNode = l.Head
		l.NextNode.Next = l.PrevNode
		l.Head = l.Head.Next
	}
	l.Head = l.PrevNode
}
