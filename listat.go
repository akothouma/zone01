package piscine

type NodeL5 struct {
	Data interface{}
	Next *NodeL
}

type List5 struct {
	Head *NodeL
	Tail *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	traversalcount := 0
	Head := l
	if pos == 0 && Head != nil {
		return l
	} else {
		for Head != nil && traversalcount != pos {
			traversalcount++
			Head = Head.Next
			if traversalcount == pos {
				l = Head
				return l
			}
		}
		return nil
	}
}
