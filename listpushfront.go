package piscine

type NodeL2 struct {
	Data interface{}
	Next *NodeL
}

type List2 struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List2, data interface{}) {
	newerNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newerNode
	} else {
		newerNode.Next = l.Head
	}
	l.Head = newerNode
}
