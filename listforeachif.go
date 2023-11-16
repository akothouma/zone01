package piscine

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	TempNode := l.Head
	for TempNode != nil {
		if cond(TempNode) {
			f(TempNode)
		}
		TempNode = TempNode.Next
	}
}
func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
