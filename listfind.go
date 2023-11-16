package piscine

func CompStr(a, b interface{}) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	var wantedvalue *interface{}
	TempNode := l.Head
	for TempNode != nil {
		compData := TempNode.Data
		if comp(compData, ref) {
			wantedvalue = &TempNode.Next.Data
		}
		TempNode = TempNode.Next
	}
	return wantedvalue
}
