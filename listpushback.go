package piscine

type NodeL8 struct {
	Data interface{}
	Next *NodeL
}

type List8 struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List8, data interface{}) {
	newerNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newerNode
	} else {
		l.Tail.Next = newerNode
	}
	l.Tail = newerNode
}
