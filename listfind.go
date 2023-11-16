package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	var wantedvalue *interface{}
	TempNode := l.Head
	for TempNode != nil {
		compData := TempNode.Data
		if comp(compData, ref) {
			wantedvalue = &TempNode.Data
		}
		TempNode = TempNode.Next
	}
	return wantedvalue
}
