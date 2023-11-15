package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	count := pos - 1
	traversalcount := 0
	Head := l
	for Head != nil && traversalcount != count {
		traversalcount++
		Head = Head.Next
		if traversalcount == count {
			l = Head
			return l
		}
	}
	return nil
}
