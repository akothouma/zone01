package piscine

type NodeL1 struct {
	Data interface{}
	Next *NodeL
}

type List1 struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
	} else {
		l.Tail.Next = newNode
	}
	l.Tail = newNode
}
