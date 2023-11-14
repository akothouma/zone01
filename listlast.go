package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
	Name string
}

func NameOfList(l *List) {
	l.Name = "list"
}

func ListLast(l *List) interface{} {
	if l.Name == "list" {
		return l.Tail.Data
	}
	return nil
}
