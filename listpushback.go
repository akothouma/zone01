package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	if l.Tail.Next == nil {
		newNode := &NodeL{Data: data}
		l.Tail.Next = newNode
		newNode.Next = nil
	}

}
