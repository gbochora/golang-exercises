package wordfreq

import (
	"bufio"
	"io"
)

func countWordsFreq(r io.Reader) map[string]int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	counts := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}
	return counts
}
