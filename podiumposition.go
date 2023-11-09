package piscine

func PodiumPosition(podium [][]string) [][]string {
	for index := 0; index < len(podium); index++ {
		podium[index] = podium[len(podium)-1-index]
	}
	return podium
}
