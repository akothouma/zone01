package piscine

func Join(strs []string, sep string) string {
	joinedstring := strs[0]

	for i := 1; i < len(strs); i++ {
		joinedstring += sep + strs[i]
	}

	return joinedstring
}
