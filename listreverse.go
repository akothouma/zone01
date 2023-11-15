package piscine

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	var value interface{}
	if l.Tail != l.Head {
		ptr := l.Tail
		value = &NodeL{Data: ptr.Data}
		ptr.Next.Next = l.Tail
	}
	value = l.Head.Data
	fmt.Println(value)
}
