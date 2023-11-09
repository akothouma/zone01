package piscine

func ReverseMenuIndex(menu []string) []string {
	var correctmenu []string
	for index := 0; index < len(menu)-1; index++ {
		correctmenu[index] = menu[len(menu)-1-index]
	}
	return correctmenu
}
