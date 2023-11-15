package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head        *NodeL
	Tail        *NodeL
	PrevNodeptr *NodeL
	NextNodeptr *NodeL
}

func ListReverse(l *List) {
	l.PrevNodeptr = nil
	l.NextNodeptr = nil
	for l.Head != nil {
		l.NextNodeptr = l.Head.Next
		l.Head.Next = l.PrevNodeptr
		l.PrevNodeptr = l.Head
		l.NextNodeptr.Next = l.PrevNodeptr
		l.Head = l.NextNodeptr
	}
	l.Head = l.PrevNodeptr
}
