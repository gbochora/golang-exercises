package duplicats

func eliminateAdjacentDup(strings []string) []string {
	index := 0
	for _, w := range strings {
		if strings[index] == w {
			continue
		}
		index++
		strings[index] = w
	}
	return strings[:index+1]
}
