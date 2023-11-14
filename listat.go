package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	count := pos - 1
	traversalcount := 0
	Head := l
	for Head != nil {
		traversalcount++
		if traversalcount != count {
			Head = Head.Next
		}
		l = Head
		return l
	}
	return nil
}
