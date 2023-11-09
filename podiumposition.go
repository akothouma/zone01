package piscine

func PodiumPosition(podium [][]string) [][]string {
	correctorder := [][]string{}
	for index := 0; index < len(podium); index++ {
		correctorder := [][]string{}
		correctorder[index] = podium[len(podium)-1-index]

	}
	return correctorder
}
